package pusher

import (
	"github.com/tomatod/statsd-push-agent/agent/metric"	
	"github.com/tomatod/statsd-push-agent/agent/config"
	"github.com/tomatod/statsd-push-agent/agent/log"

	"net"
	"fmt"
	"strconv"
	"time"
)

type Pusher struct {
	con    				net.Conn
	config 				*config.Config
}

func NewPusher(conf *config.Config) (*Pusher, error) {
	conn, err := net.Dial("udp4", conf.Server.Address)
	if err != nil {
		return nil, err
	}
	return &Pusher {
		con: conn,
		config: conf,
	}, nil
}

func (p *Pusher) StartLoop(isStopped chan bool) {
	defer log.L.Sync()
	ticker := time.NewTicker(time.Second * time.Duration(p.config.Server.Period))
	defer ticker.Stop()
	loop: for {
		select {
		case <- isStopped:
			log.L.Info("put-metric loop thread received stop signal from main thread.")
			break loop
		case <- ticker.C:
			p.PushMetrics()
		}
	}
	isStopped <- true
	close(isStopped)
	return
}

func (p *Pusher) PushMetrics() {
	for _, m := range p.config.Metrics {
  	data, err := createStatsdMetricData(p.config.Server.Prefix, m)
		if err != nil {
			log.L.Warnf("MetricsPushError: metrics: %s  err: %s", m.Name, err)
			continue
		}
  	_, err = p.con.Write(data)
		if err != nil {
			log.L.Warnf("MetricsPushError: metrics: %s  err: %s", m.Name, err)
			continue
		}
	}
}

func createStatsdMetricData(prefix string, mtrc *metric.Metric) ([]byte, error) {
	v, err := mtrc.GetBlendedMetricValue()
	if err != nil {
		return nil, err
	}
	str := fmt.Sprintf("%s.%s:%s%s|%s", 
	 	prefix,
		mtrc.Name,
		mtrc.Operator,
		strconv.FormatFloat(v, 'f', -1, 64),
		mtrc.Type,
	)
	if mtrc.Rate != 0 {
		str += fmt.Sprintf("|@%s", strconv.FormatFloat(mtrc.Rate, 'f', -1, 64))
	}
	return []byte(str), nil
}