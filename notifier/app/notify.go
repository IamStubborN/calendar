package app

import (
	"github.com/IamStubborN/calendar/notifier/pkg/broker"
	"github.com/IamStubborN/calendar/notifier/pkg/logger"
	"github.com/IamStubborN/calendar/notifier/pkg/notify/service"
	"github.com/IamStubborN/calendar/notifier/worker"
)

func initializeNotifyService(logger logger.UseCase, br broker.Repository) worker.Worker {
	return service.NewNotifyService(logger, br)
}
