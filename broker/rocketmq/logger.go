package rocketmq

import "github.com/go-kratos/kratos/v2/log"

type logger struct {
	log.Logger
}

func (l logger) Debug(msg string, fields map[string]interface{}) {
	l.Logger.Log(log.LevelDebug, msg)
}

func (l logger) Info(msg string, fields map[string]interface{}) {
	l.Logger.Log(log.LevelInfo, msg)
}

func (l logger) Warning(msg string, fields map[string]interface{}) {
	l.Logger.Log(log.LevelWarn, msg)
}

func (l logger) Error(msg string, fields map[string]interface{}) {
	l.Logger.Log(log.LevelError, msg)
}

func (l logger) Fatal(msg string, fields map[string]interface{}) {
	l.Logger.Log(log.LevelFatal, msg)
}

func (l logger) Level(level string) {
}

func (l logger) OutputPath(path string) (err error) {
	return nil
}
