package metric

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMetricInitialize(t *testing.T) {
  actual := Metric {
    Blend: "/100",
  }
  e := actual.Initialize()
  if e != nil {
    t.Fatal(e)
  }
  expect := Metric {
    Blend: "/100",
    blendValue: 100,
    blendOpe: "/",
  }
  assert.Equal(t, expect, actual, "Not equal.")
}

func TestGetStrValueByFileMethod(t *testing.T) {
  m := Metric {
    Method: "file", 
    Param: "/dev/null",
  }
  s, e := m.getStrValueByFileMethod()
  if e != nil {
    t.Fatal(e)
  }
  assert.Equal(t, "", s, "Not equal.")
}

func TestGetStrValueByCmdMethod(t *testing.T) {
  m := Metric {
    Method: "cmd", 
    Param: "echo 1.5",
  }
  s, e := m.getStrValueByCmdMethod()
  if e != nil {
    t.Fatal(e)
  }
  assert.Equal(t, "1.5\n", s, "Not equal.")
}

func TestGetMetricValue(t *testing.T) {
  m := Metric {
    Method: "cmd", 
    Param: "echo 1.2",
  }
  v, e := m.getMetricValue()
  if e != nil {
    t.Fatal(e)
  }
  assert.Equal(t, 1.2, v, "Not equal.")
}

func TestGetBlendedValue(t *testing.T) {
  m := Metric {
    Method: "cmd", 
    Param: "echo 125",
    Blend: "/100",
  }
  e := m.Initialize()
  if e != nil {
    t.Fatal(e)
  }

  v, e := m.GetBlendedMetricValue()
  if e != nil {
    t.Fatal(e)
  }
  assert.Equal(t, 1.25, v, "Not equal.")
}