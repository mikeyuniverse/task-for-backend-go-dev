package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logg interface {
	Fatal(err string)
	Error(err string)
	Warn(err string)
	Info(err string)
}

type Logger struct {
	log *logrus.Logger
}

func New() *Logger {
	log := logrus.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
	return &Logger{log: log}
}

func (l *Logger) Fatal(err string) {
	l.log.Fatal(err)
}

func (l *Logger) Error(err string) {
	l.log.Error(err)
}

func (l *Logger) Warn(err string) {
	l.log.Warn(err)
}

func (l *Logger) Info(err string) {
	l.log.Info(err)
}
