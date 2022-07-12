package config

import (
	"github.com/tomatod/statsd-push-agent/agent/metric"
  "go.uber.org/zap"

  "testing"
  "github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
  actual, err := ReadConfig("config_test.yaml")
	if err != nil {
		t.Fatal(err)
	}
  expect := Config {
    Server {
      Address: "127.0.0.1:8125", 
      Prefix: "statsd", 
      Period: 5,
    },
    &zap.Config {},
    []*metric.Metric {
      &metric.Metric {
        Name: "hoge",
        Method: "cmd", 
        Param: "echo 1",
        Type: "s",
        Rate: 1,
        Blend: "",
      },
      &metric.Metric {
        Name: "bar",
        Method: "file", 
        Param: "/proc/sys/vm/overcommit_ratio",
        Type: "g",
        Operator: "+",
        Rate: 0.5,
        Blend: "/10",
      },
    },
  }
  assert.Equal(t, &expect, actual, "Not equal.")
}