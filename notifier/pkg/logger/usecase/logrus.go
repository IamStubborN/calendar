package usecase

import (
	"os"

	"github.com/IamStubborN/calendar/notifier/pkg/logger"
	"github.com/sirupsen/logrus"
)

type log struct {
	logger *logrus.Logger
}

func NewLoggerLogrus(level string) (logger.UseCase, error) {
	l := logrus.New()
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		l.Fatalln(err)
		return nil, err
	}

	l.SetLevel(lvl)
	l.SetOutput(os.Stdout)

	return &log{
		logger: l,
	}, nil
}

func (l *log) Info(data ...interface{}) {
	l.logger.Infoln(data...)
}

func (l *log) Warn(data ...interface{}) {
	l.logger.Warnln(data...)
}

func (l *log) Fatal(data ...interface{}) {
	l.logger.Fatalln(data...)
}

func (l *log) WithFields(level string, data map[string]interface{}, msg ...interface{}) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}

	l.logger.WithFields(data).Log(lvl, msg...)
}
