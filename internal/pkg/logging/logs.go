package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Log struct {
	log logrus.Logger
}

func (l *Log) Init() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Не удалось открыть файл логов, используется стандартный stderr")
	}
	return log
}
