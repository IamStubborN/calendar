package grpc

import (
	"context"
	"github.com/IamStubborN/calendar/pkg/event"

	"github.com/IamStubborN/calendar/models"
	"github.com/IamStubborN/calendar/pkg/event/delivery/grpc/event_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
)

type Server struct {
	Storage event.Repository
}

func NewEventGRPCServer(storage event.Repository) *Server {
	return &Server{
		Storage: storage,
	}
}

func (s *Server) Create(ctx context.Context, req *event_grpc.CreateRequest) (*event_grpc.CreateResponse, error) {
	ev := &models.Event{
		ID:          req.Event.ID,
		Name:        req.Event.Name,
		Description: req.Event.Description,
		Date:        req.Event.Date,
	}

	ev, err := s.Storage.Create(ctx, ev)
	if err != nil {
		return nil, err
	}

	gEvent := &event_grpc.Event{
		ID:          ev.ID,
		Name:        ev.Name,
		Description: ev.Description,
		Date:        ev.Date,
	}

	return &event_grpc.CreateResponse{
		Event: gEvent,
	}, nil
}

func (s *Server) Read(ctx context.Context, req *event_grpc.ReadRequest) (*event_grpc.ReadResponse, error) {
	ev, err := s.Storage.Read(ctx, req.Event_ID)
	if err != nil {
		return nil, err
	}

	gEvent := &event_grpc.Event{
		ID:          ev.ID,
		Name:        ev.Name,
		Description: ev.Description,
		Date:        ev.Date,
	}

	return &event_grpc.ReadResponse{Event: gEvent}, nil
}

func (s *Server) Update(ctx context.Context, req *event_grpc.UpdateRequest) (*event_grpc.UpdateResponse, error) {
	ev := &models.Event{
		ID:          req.Event.ID,
		Name:        req.Event.Name,
		Description: req.Event.Description,
		Date:        req.Event.Date,
	}

	updated, err := s.Storage.Update(ctx, ev)
	if err != nil {
		return nil, err
	}

	return &event_grpc.UpdateResponse{
		Updated: updated,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *event_grpc.DeleteRequest) (*event_grpc.DeleteResponse, error) {
	deleted, err := s.Storage.Delete(ctx, req.Event_ID)
	if err != nil {
		return nil, err
	}

	return &event_grpc.DeleteResponse{
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
	event_grpc.RegisterEventServiceServer(gServer, s)

	go func() {
		<-ctx.Done()
		gServer.GracefulStop()
	}()

	if err := gServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
