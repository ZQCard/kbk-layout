package loggerHelper

import (
	ktrus "github.com/go-kratos/kratos/contrib/log/logrus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sirupsen/logrus"
)

func NewLogger() log.Logger {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	return ktrus.NewLogger(logrus.StandardLogger())
}
