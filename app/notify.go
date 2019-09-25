package app

import (
	"github.com/IamStubborN/calendar/pkg/broker"
	"github.com/IamStubborN/calendar/pkg/logger"
	"github.com/IamStubborN/calendar/pkg/notify/service"
	"github.com/IamStubborN/calendar/worker"
)

func initializeNotifyService(logger logger.Repository, br broker.Repository) worker.Worker {
	return service.NewNotifyService(logger, br)
}
