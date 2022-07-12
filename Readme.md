# statsd-push-agent
statsd-push-agent is agent application for pushing metrics to StatsD server (StatsD, CloudWatch Agent, Datadog Agent and so on). In other word, statsd-push-agent act as StatsD client.

## Install and start
### Linux
1. Download zip from [Git Hub release]("/released"). 
2. Run next command lines for installing.
```
unzip <downloaded zip file>
cd statsd-push-agent
chmod +x ./install.sh
sudo ./install.sh
```
3. Rewrite /opt/statsd-push-agent/config.yaml.
4. Run "sudo systemctl start statsd-push-agent.service" for starting.

### Windows or Mac
1. Download zip installer from [Git Hub release]("/released"). 
2. Unzip downloaded zip file
3. Move statsd-push-agent directory to any location.
4. Rewrite config.yaml and run statsd-push-agent binary in statsd-push-agent directory.

## Configuration
statsd-push-agent works with a simple yaml config file.

``` yaml
# information of destination StatsD server
server:
  address: "<host>:<port>" #require
  prefix:  "<top-level prefix of all metrics>" #option
  period:   <push intervarl (second)> #require

# logging configuration
logging:
  level: "<log level>" # require
  outputPaths:
    - "<file path> or stderr or stdout"
  errorOutputPaths:
    - "<file path> or stderr or stdout"

# information of metrics you want to push.
metrics:
  - name:     "<metrics-level prefix>" #require
    method:   "<method of getting data of a metric>" #require
    param:    "<parameter for above method>" #require
    type:     "<metrics type of StatsD>" #require
    operator: "<operator (+/-) of a metric. some type use.>" #option
    rate:      <rate of StatsD. some type use.> #option
    blend:    "<four arithmetic operations　to raw value>" #option
```

| parameter | supplement |
| --- | --- |
| metrics[].method | Method to get metric data. You can specify only "cmd" or "file". |
| metrics[].param | In metrics[].method, if you choose "cmd", metrics[].param is shell command to get metric data. If you choose "file", metrics[].method is file path having data. |
| metrics[].type | You can specify any type of StatsD with string. For example, gauge is "g", timing (timer) is "ms", count is "c" and so on. You should confirm the specification of StatsD server which you use. |
| metrics[].blend | You can adjust the metric values, ​​using a simple four arithmetic operation (+, -, *, /). For example, if this parameter is "/100" and raw metric data is 50, pushed metric data become 50 / 100 = 0.5. Similarly, "+10" and 1 is 11, "*20" and 0.5 is 10, "-5" and 12.1 is 7.1. |

### Example
``` yaml
server:
  address: "127.0.0.1:8125"
  prefix: "sample"
  period: 60

logger:
  level: "warn"
  outputPaths:
    - "/opt/statsd-push-agent/statsd-push-agent.log"
  errorOutputPaths:
    - "/opt/statsd-push-agent/statsd-push-agent.log"
    - "stderr"

metrics:
  # Statsd packet: sample.echo_one:1|s
  - name: "echo_one"
    method: "cmd"
    param: "echo 1"
    type: "s"

  # Statsd packet: sample.cpu_temp:50|g
  - name: "cpu_temp"
    method: "file"
    param: "/sys/class/thermal/thermal_zone0/temp" # for example 50000
    type: "g"
    blend: "/1000"
```