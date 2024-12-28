package helper

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
	}
	return logger
}

func StringQueryLike(value string) string {
	return "%" + strings.ToLower(value) + "%"
}

func StringToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	PanicIfError(err)
	return intValue
}

func StringToInt64(value string) int64 {
	intValue, err := strconv.ParseInt(value, 10, 64)
	PanicIfError(err)
	return intValue
}
