package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func New() (log *logrus.Logger) {
	log = logrus.New()
	log.Out = os.Stdout
	return
}
