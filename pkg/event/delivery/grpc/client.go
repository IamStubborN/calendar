package grpc

import (
	"context"

	"github.com/IamStubborN/calendar/pkg/event/delivery/grpc/event_grpc"
	"github.com/IamStubborN/calendar/pkg/logger"

	"google.golang.org/grpc"
)

type Client struct {
	logger logger.Repository
	gc     event_grpc.EventServiceClient
}

func NewEventGRPCClient(logger logger.Repository) (*Client, error) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		logger: logger,
		gc:     event_grpc.NewEventServiceClient(cc),
	}, nil
}

func (c *Client) Run(ctx context.Context) error {
	ev := &event_grpc.Event{
		ID:          0,
		Name:        "First event",
		Description: "This is my first event with grpc",
		Date:        "2019-08-22",
	}

	ev, err := c.Create(ctx, ev)
	if err != nil {
		return err
	}

	c.logger.WithFields("info", map[string]interface{}{
		"event": ev.ID,
		"name":  ev.Name,
		"date":  ev.Date,
		"desc":  ev.Description,
	}, "event grpc client: create")

	ev, err = c.Read(ctx, ev.ID)
	if err != nil {
		return err
	}

	c.logger.WithFields("info", map[string]interface{}{
		"event": ev.ID,
		"name":  ev.Name,
		"date":  ev.Date,
		"desc":  ev.Description,
	}, "event grpc client: read")

	ev.Name = "Updated First Event"

	updated, err := c.Update(ctx, ev)
	if err != nil {
		return err
	}

	c.logger.WithFields("info", map[string]interface{}{
		"update": updated,
	}, "event grpc client: update")

	deleted, err := c.Delete(ctx, ev.ID)
	if err != nil {
		return err
	}

	c.logger.WithFields("info", map[string]interface{}{
		"delete": deleted,
	}, "event grpc client: delete")

	return nil
}

func (c *Client) Create(ctx context.Context, ev *event_grpc.Event) (*event_grpc.Event, error) {
	resp, err := c.gc.Create(ctx, &event_grpc.CreateRequest{
		Event: ev,
	})

	if err != nil {
		return nil, err
	}

	return resp.Event, nil
}

func (c *Client) Read(ctx context.Context, eventID uint64) (*event_grpc.Event, error) {
	resp, err := c.gc.Read(ctx, &event_grpc.ReadRequest{
		Event_ID: eventID,
	})

	if err != nil {
		return nil, err
	}

	return resp.Event, nil
}

func (c *Client) Update(ctx context.Context, ev *event_grpc.Event) (bool, error) {
	resp, err := c.gc.Update(ctx, &event_grpc.UpdateRequest{
		Event: ev,
	})
	if err != nil {
		return false, err
	}

	return resp.Updated, nil
}

func (c *Client) Delete(ctx context.Context, eventID uint64) (bool, error) {
	resp, err := c.gc.Delete(ctx, &event_grpc.DeleteRequest{
		Event_ID: eventID,
	})
	if err != nil {
		return false, err
	}

	return resp.Deleted, nil
}
