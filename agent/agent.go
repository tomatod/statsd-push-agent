package main

import (
	"github.com/tomatod/statsd-push-agent/agent/config"
	"github.com/tomatod/statsd-push-agent/agent/pusher"
	"github.com/tomatod/statsd-push-agent/agent/log"

	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	confPath *string
)

func init() {
	confPath = flag.String("c", "/opt/statsd-push-agent/config.yaml", "config file path")
	flag.Parse()
}

func main() {
	defer	func() {
		log.OneshotLogger().Info("statsd-push-agent is stopped.")
	}()
	
	cnf, err := config.InitializeConfig(*confPath)
  if err != nil {
	  log.OneshotLogger().Errorf("ConfigInitError: %s", err)
		return 
  }

	err = log.LogInit(cnf.Logging)
	if err != nil {
		log.OneshotLogger().Errorf("LogInitError: %s", err)
		return
	}
	defer	func() {
		log.L.Sync()
	}()

	pusher, err := pusher.NewPusher(cnf)
	if err != nil {
	  log.L.Errorf("PusherCreateError: %s", err)
		return
  }
	
	isStopped := make(chan bool)
	go pusher.StartLoop(isStopped)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	select {
	case <- isStopped:
		log.L.Errorf("LoopStoppedError: put-metric loop is unexpectedlly stopped.")
	case s := <- sig:
		log.L.Infof("statsd-push-agent recieved OS signal (%s).", s.String())
		isStopped <- true
		<- isStopped
	}
}