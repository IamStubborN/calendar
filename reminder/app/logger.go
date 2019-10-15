package app

import (
	"github.com/IamStubborN/calendar/reminder/config"
	"github.com/IamStubborN/calendar/reminder/pkg/logger"
	"github.com/IamStubborN/calendar/reminder/pkg/logger/usecase"
	"github.com/sirupsen/logrus"
)

func initializeLogger(cfg *config.Config) logger.UseCase {
	log, err := usecase.NewLoggerLogrus(cfg.Logger.Level)
	if err != nil {
		logrus.Fatalln(err)
	}

	return log
}
