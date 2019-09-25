package service

import (
	"context"
	"time"

	"github.com/IamStubborN/calendar/pkg/broker"
	"github.com/IamStubborN/calendar/pkg/logger"
	"github.com/IamStubborN/calendar/worker"
)

type notifyService struct {
	logger logger.Repository
	broker broker.Repository
}

func NewNotifyService(logger logger.Repository, br broker.Repository) worker.Worker {
	return &notifyService{
		logger: logger,
		broker: br,
	}
}

func (ns notifyService) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			ns.logger.Info("notify service closed")
			return nil
		case <-time.Tick(time.Second):
			ns.logger.Info(ns.broker.Receive("remind"))
		}
	}
}
