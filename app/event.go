package app

import (
	"github.com/IamStubborN/calendar/config"
	"github.com/IamStubborN/calendar/pkg/event"
	"github.com/IamStubborN/calendar/pkg/event/repository"
	"github.com/IamStubborN/calendar/pkg/event/service"
	"github.com/IamStubborN/calendar/pkg/logger"
	"github.com/IamStubborN/calendar/worker"
	"github.com/jmoiron/sqlx"
)

func initializeEventService(logger logger.Repository, er event.Repository) worker.Worker {
	EService, err := service.NewEventService(logger, er)
	if err != nil {
		logger.Fatal(err)
	}

	return EService
}

func initializeEventRepository(cfg *config.Config, pool *sqlx.DB) event.Repository {
	var storage event.Repository

	switch cfg.Storage.Provider {
	case "postgres":
		storage = repository.NewEventRepositoryPSQL(pool)
	case "cache":
		storage = repository.NewEventRepositoryCache()
	}

	return storage
}
