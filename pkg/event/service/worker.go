package service

import (
	"context"

	"github.com/IamStubborN/calendar/pkg/event"

	"github.com/IamStubborN/calendar/pkg/logger"

	"github.com/IamStubborN/calendar/pkg/event/delivery/grpc"
	"github.com/IamStubborN/calendar/worker"
)

type eventService struct {
	logger logger.Repository
	client *grpc.Client
	server *grpc.Server
}

func NewEventService(logger logger.Repository, storage event.Repository) (worker.Worker, error) {
	client, err := grpc.NewEventGRPCClient(logger)
	if err != nil {
		logger.WithFields("fatal", map[string]interface{}{
			"eventService": "error",
			"type":         err,
		}, "new event worker error")
	}

	return &eventService{
		logger: logger,
		client: client,
		server: grpc.NewEventGRPCServer(storage),
	}, nil
}

func (e eventService) Run(ctx context.Context) error {
	var err error

	go func() {
		if err = e.server.Run(ctx); err != nil {
			e.logger.Warn("grpc server", err)
			return
		}
	}()

	if err = e.client.Run(ctx); err != nil {
		e.logger.Warn("grpc client", err)
		return err
	}

	return err
}
