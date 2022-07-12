package pusher

import (
	"github.com/tomatod/statsd-push-agent/agent/metric"	

	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateStatsdMetricData1(t *testing.T) {
	prefix := "test"
	m := &metric.Metric {
		Name: "statsd",
		Method: "cmd",
		Param: "echo 120",
		Type: "ms",
		Rate: 0.1,
		Blend: "/1000",
	}
	err := m.Initialize()
	if err != nil {
		t.Fatal(err)
	}
	actual, err := createStatsdMetricData(prefix, m)
	if err != nil {
		t.Fatal(err)
	}
	expect := []byte("test.statsd:0.12|ms|@0.1")
	assert.Equal(t, expect, actual, "Not equal.")
}

func TestCreateStatsdMetricData2(t *testing.T) {
	prefix := "test"
	m := &metric.Metric {
		Name: "statsd",
		Method: "cmd",
		Param: "echo 5.5",
		Type: "g",
		Operator: "+",
	}
	err := m.Initialize()
	if err != nil {
		t.Fatal(err)
	}
	actual, err := createStatsdMetricData(prefix, m)
	if err != nil {
		t.Fatal(err)
	}
	expect := []byte("test.statsd:+5.5|g")
	assert.Equal(t, expect, actual, "Not equal.")
}