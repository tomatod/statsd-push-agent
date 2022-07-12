package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.SugaredLogger
func LogInit(config *zap.Config) error {
	cnf := zap.NewProductionConfig()
	cnf.Level = config.Level
	cnf.Encoding = "console"
	cnf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if config.OutputPaths != nil {
		cnf.OutputPaths = config.OutputPaths
	}
	if config.ErrorOutputPaths != nil {
		cnf.ErrorOutputPaths = config.ErrorOutputPaths
	}
	gen, err := cnf.Build()
	if err != nil {
		return err
	}
	L = gen.Sugar()
	return nil
}

func OneshotLogger() *zap.SugaredLogger {
	l, _ := zap.NewDevelopment()
	return l.Sugar()
}