package app

import (
	"github.com/IamStubborN/calendar/config"
	"github.com/IamStubborN/calendar/pkg/broker"
	"github.com/IamStubborN/calendar/pkg/logger"
	"github.com/IamStubborN/calendar/pkg/remind"
	"github.com/IamStubborN/calendar/pkg/remind/repository"
	"github.com/IamStubborN/calendar/pkg/remind/service"
	"github.com/IamStubborN/calendar/worker"
	"github.com/jmoiron/sqlx"
	"time"
)

func initializeRemindService(freq time.Duration, logger logger.Repository, rr remind.Repository, br broker.Repository) worker.Worker {
	RService, err := service.NewRemindService(freq, logger, rr, br)
	if err != nil {
		logger.Fatal(err)
	}

	return RService
}

func initializeRemindRepository(cfg *config.Config, pool *sqlx.DB) remind.Repository {
	var storage remind.Repository
	switch cfg.Storage.Provider {
	case "postgres":
		storage = repository.NewRemindRepositoryPSQL(pool)
	}

	return storage
}
