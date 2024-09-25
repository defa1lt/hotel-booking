package logger

import (
	"github.com/sirupsen/logrus"
	"strings"
)

func NewLogger(level string) *logrus.Logger {
	log := logrus.New()
	lvl, err := logrus.ParseLevel(strings.ToLower(level))
	if err != nil {
		lvl = logrus.InfoLevel
	}
	log.SetLevel(lvl)
	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}
