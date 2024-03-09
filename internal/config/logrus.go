package config

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(6)
	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}
