package grpc

import (
	"context"

	"github.com/IamStubborN/calendar/api/pkg/event/delivery/grpc/entries"
	"github.com/IamStubborN/calendar/api/pkg/logger"

	"google.golang.org/grpc"
)

type Client struct {
	logger logger.UseCase
	gc     entries.EventServiceClient
}

func NewEventGRPCClient(logger logger.UseCase) (*Client, error) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		logger: logger,
		gc:     entries.NewEventServiceClient(cc),
	}, nil
}

func (c *Client) Run(ctx context.Context) error {
	ev := &entries.Event{
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

func (c *Client) Create(ctx context.Context, ev *entries.Event) (*entries.Event, error) {
	resp, err := c.gc.Create(ctx, &entries.CreateRequest{
		Event: ev,
	})

	if err != nil {
		return nil, err
	}

	return resp.Event, nil
}

func (c *Client) Read(ctx context.Context, eventID uint64) (*entries.Event, error) {
	resp, err := c.gc.Read(ctx, &entries.ReadRequest{
		Event_ID: eventID,
	})

	if err != nil {
		return nil, err
	}

	return resp.Event, nil
}

func (c *Client) Update(ctx context.Context, ev *entries.Event) (bool, error) {
	resp, err := c.gc.Update(ctx, &entries.UpdateRequest{
		Event: ev,
	})
	if err != nil {
		return false, err
	}

	return resp.Updated, nil
}

func (c *Client) Delete(ctx context.Context, eventID uint64) (bool, error) {
	resp, err := c.gc.Delete(ctx, &entries.DeleteRequest{
		Event_ID: eventID,
	})
	if err != nil {
		return false, err
	}

	return resp.Deleted, nil
}
