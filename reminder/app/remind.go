package app

import (
	"time"

	"github.com/IamStubborN/calendar/reminder/config"
	"github.com/IamStubborN/calendar/reminder/pkg/broker"
	"github.com/IamStubborN/calendar/reminder/pkg/logger"
	"github.com/IamStubborN/calendar/reminder/pkg/remind"
	"github.com/IamStubborN/calendar/reminder/pkg/remind/repository"
	"github.com/IamStubborN/calendar/reminder/pkg/remind/service"
	"github.com/IamStubborN/calendar/reminder/worker"
	"github.com/jmoiron/sqlx"
)

func initializeRemindService(
	freq time.Duration,
	logger logger.UseCase,
	rr remind.Repository,
	br broker.Repository) worker.Worker {
	RService, err := service.NewRemindService(freq, logger, rr, br)
	if err != nil {
		logger.Fatal(err)
	}

	return RService
}

func initializeRemindRepository(cfg *config.Config, pool *sqlx.DB) remind.Repository {
	var storage remind.Repository
	if cfg.Storage.Provider == "postgres" {
		storage = repository.NewRemindRepositoryPSQL(pool)
	}

	return storage
}
