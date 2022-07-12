package config

import (
	"github.com/tomatod/statsd-push-agent/agent/metric"

  "os"
  "gopkg.in/yaml.v2"
	"go.uber.org/zap"
)

type Config struct {
  Server  Server   `yaml:"server"`
  Logging  *zap.Config `yaml:"logging"`
  Metrics []*metric.Metric `yaml:"metrics"`
}

type Server struct {
  Address string `yaml:"address"`
  Prefix  string `yaml:"prefix"`
  Period  int    `yaml:"period"`
}

func InitializeConfig(confpath string) (*Config, error) {
  cnf, err := ReadConfig(confpath)
  if err != nil {
    return nil, err
  }
  if err = cnf.InitializeAll(); err != nil {
		return nil, err
	}
  return cnf, nil
}

func ReadConfig(confpath string) (*Config, error) {
  cnf := Config{Server{}, &zap.Config{}, []*metric.Metric{}}
  bytes, err := os.ReadFile(confpath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bytes, &cnf)
	if err != nil {
		return nil, err
	}
  return &cnf, nil
}

func (c *Config) InitializeAll() error {
  for _, m := range c.Metrics {
    if err := m.Initialize(); err != nil {
      return err
    }
  }
  return nil
}