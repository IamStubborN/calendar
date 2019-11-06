package grpc

import (
	"context"

	"github.com/IamStubborN/calendar/api/pkg/event"
	"github.com/IamStubborN/calendar/api/pkg/logger"

	"github.com/IamStubborN/calendar/api/models"
	"github.com/IamStubborN/calendar/api/pkg/event/delivery/grpc/entries"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
)

type Server struct {
	logger          logger.UseCase
	eventRepository event.Repository
}

func NewEventGRPCServer(logger logger.UseCase, er event.Repository) *Server {
	return &Server{
		logger:          logger,
		eventRepository: er,
	}
}

func (s *Server) Create(ctx context.Context, req *entries.CreateRequest) (*entries.CreateResponse, error) {
	ev := &models.Event{
		ID:          req.Event.ID,
		Name:        req.Event.Name,
		Description: req.Event.Description,
		Date:        req.Event.Date,
	}

	ev, err := s.eventRepository.Create(ctx, ev)
	if err != nil {
		return nil, err
	}

	gEvent := &entries.Event{
		ID:          ev.ID,
		Name:        ev.Name,
		Description: ev.Description,
		Date:        ev.Date,
	}

	return &entries.CreateResponse{
		Event: gEvent,
	}, nil
}

func (s *Server) Read(ctx context.Context, req *entries.ReadRequest) (*entries.ReadResponse, error) {
	ev, err := s.eventRepository.Read(ctx, req.Event_ID)
	if err != nil {
		return nil, err
	}

	gEvent := &entries.Event{
		ID:          ev.ID,
		Name:        ev.Name,
		Description: ev.Description,
		Date:        ev.Date,
	}

	return &entries.ReadResponse{Event: gEvent}, nil
}

func (s *Server) Update(ctx context.Context, req *entries.UpdateRequest) (*entries.UpdateResponse, error) {
	ev := &models.Event{
		ID:          req.Event.ID,
		Name:        req.Event.Name,
		Description: req.Event.Description,
		Date:        req.Event.Date,
	}

	updated, err := s.eventRepository.Update(ctx, ev)
	if err != nil {
		return nil, err
	}

	return &entries.UpdateResponse{
		Updated: updated,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *entries.DeleteRequest) (*entries.DeleteResponse, error) {
	deleted, err := s.eventRepository.Delete(ctx, req.Event_ID)
	if err != nil {
		return nil, err
	}

	return &entries.DeleteResponse{
		Deleted: deleted,
	}, nil
}

func (s *Server) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		return err
	}

	gServer := grpc.NewServer()
	reflection.Register(gServer)
	entries.RegisterEventServiceServer(gServer, s)

	go func() {
		<-ctx.Done()
		s.logger.Info("event service closed")
		gServer.GracefulStop()
	}()

	if err := gServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
