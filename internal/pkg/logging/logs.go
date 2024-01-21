package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Log struct {
	Log *logrus.Logger
}

func NewLog() *Log {
	return &Log{
		Log: logrus.New(),
	}
}
func (l *Log) Init() *logrus.Logger {
	l.Log.SetLevel(logrus.DebugLevel)
	l.Log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		l.Log.SetOutput(file)
	} else {
		l.Log.Info("Не удалось открыть файл логов, используется стандартный stderr")
	}
	return l.Log
}
