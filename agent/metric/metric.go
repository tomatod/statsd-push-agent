package metric

import (
  "io/ioutil"
  "os/exec"
  "strings"
  "strconv"
  "regexp"
  "fmt"
)

type Metric struct {
  Name       string  `yaml:"name"`
  Method     string  `yaml:"method"`
  Param      string  `yaml:"param"`
  Type       string  `yaml:"type"`
  Operator   string  `yaml:"operator"`
  Rate       float64 `yaml:"rate"`
  Blend      string  `yaml:"blend"`
  blendValue float64 
  blendOpe   string
}

// Before being pushed, Metric must be initialized by this func.
// This is normally called in Config.Initialize() after config file is read.
func (m *Metric) Initialize() error {
  err := m.iniBlendParams()
  return err
}

// Blend value is used to adjust measured value.
// For example, if measured value: 5, blend value: +2, pushed value: 7.
// Supports only simple four arithmetic operations (+,-,*,/).
var blendRe = regexp.MustCompile(`^[\+\-\*\/][0-9]+\.*[0-9]*$`)
func (m *Metric) iniBlendParams() error {
  if m.Blend == "" {
    return nil
  }
  if !blendRe.MatchString(m.Blend) {
    return fmt.Errorf("Blend parameter '%s' in %s: %s is invalid.", m.Blend, m.Method, m.Param)
  }
  var err error
  m.blendOpe = m.Blend[0:1]
  m.blendValue, err = strconv.ParseFloat(m.Blend[1:], 64)
  return err
}

func (m *Metric) GetBlendedMetricValue() (float64, error) {
  v, err := m.getMetricValue()
  if err != nil || m.blendOpe == "" {
    return v, err
  }
  switch m.blendOpe {
  case "+": return v + m.blendValue, nil
  case "-": return v - m.blendValue, nil
  case "*": return v * m.blendValue, nil
  case "/": return v / m.blendValue, nil
  }
  return 0, fmt.Errorf("Operator '%s' for blending is invalid.", m.blendOpe)
}

func (m *Metric) getMetricValue() (float64, error) {
  var data string
  var err error
  switch m.Method {
  case "file": data, err = m.getStrValueByFileMethod()
  case "cmd" : data, err = m.getStrValueByCmdMethod()
  }
  if err != nil {
    return 0, err
  }
  data = strings.Replace(data, " ", "", -1)
  data = strings.Replace(data, "\n", "", -1)
  data = strings.Replace(data, "\r", "", -1)
  return strconv.ParseFloat(data, 64)
}

func (m *Metric) getStrValueByFileMethod() (string, error) {
  b, err := ioutil.ReadFile(m.Param)
  if err != nil {
    return "", err
  }
  return string(b), nil
}

func (m *Metric) getStrValueByCmdMethod() (string, error) {
  args := strings.Split(m.Param, " ")
  if len(args) >= 2 {
    b, err := exec.Command(args[0], args[1:]...).Output()
    return string(b), err
  }
  b, err := exec.Command(args[0]).Output()
  return string(b), err
}