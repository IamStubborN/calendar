package app

import (
	"github.com/IamStubborN/calendar/api/config"
	"github.com/IamStubborN/calendar/api/pkg/event"
	"github.com/IamStubborN/calendar/api/pkg/event/repository"
	"github.com/IamStubborN/calendar/api/pkg/event/service"
	"github.com/IamStubborN/calendar/api/pkg/logger"
	"github.com/IamStubborN/calendar/api/worker"
	"github.com/jmoiron/sqlx"
)

func initializeEventService(logger logger.UseCase, er event.Repository) worker.Worker {
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
