package helper

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
	}
	return logger
}
