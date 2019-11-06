package app

import (
	"github.com/IamStubborN/calendar/notifier/config"
	"github.com/IamStubborN/calendar/notifier/pkg/logger"
	"github.com/IamStubborN/calendar/notifier/pkg/logger/usecase"
	"github.com/sirupsen/logrus"
)

func initializeLogger(cfg *config.Config) logger.UseCase {
	log, err := usecase.NewLoggerLogrus(cfg.Logger.Level)
	if err != nil {
		logrus.Fatalln(err)
	}

	return log
}
