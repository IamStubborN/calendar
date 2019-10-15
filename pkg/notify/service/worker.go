package service

import (
	"context"

	"github.com/IamStubborN/calendar/pkg/broker"
	"github.com/IamStubborN/calendar/pkg/logger"
	"github.com/IamStubborN/calendar/worker"
)

type notifyService struct {
	logger logger.UseCase
	broker broker.Repository
}

func NewNotifyService(logger logger.UseCase, br broker.Repository) worker.Worker {
	return &notifyService{
		logger: logger,
		broker: br,
	}
}

func (ns *notifyService) Run(ctx context.Context) error {
	dataCh, err := ns.broker.Receive(ctx, "remind")
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			ns.logger.Info("notify service closed")
			return nil
		case data := <-dataCh:
			ns.logger.WithFields("info", map[string]interface{}{
				"service": "notify",
				"bytes":   len(data),
			}, "successful consumed from broker")
		}
	}
}
