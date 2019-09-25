package app

import (
	"github.com/IamStubborN/calendar/config"
	"github.com/IamStubborN/calendar/pkg/logger"
	"github.com/IamStubborN/calendar/pkg/logger/repository"
	"github.com/sirupsen/logrus"
)

func initializeLogger(cfg *config.Config) logger.Repository {
	log, err := repository.NewLoggerLogrus(cfg.Logger.Level)
	if err != nil {
		logrus.Fatalln(err)
	}

	return log
}
