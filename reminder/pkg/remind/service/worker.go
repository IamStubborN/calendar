package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/IamStubborN/calendar/reminder/pkg/broker"
	"github.com/IamStubborN/calendar/reminder/pkg/logger"
	"github.com/IamStubborN/calendar/reminder/pkg/remind"
	"github.com/IamStubborN/calendar/reminder/worker"
)

type remindService struct {
	freq   time.Duration
	logger logger.UseCase
	remind remind.Repository
	broker broker.Repository
}

func NewRemindService(
	freq time.Duration,
	logger logger.UseCase,
	rr remind.Repository,
	br broker.Repository) (worker.Worker, error) {
	return &remindService{
		freq:   freq,
		logger: logger,
		remind: rr,
		broker: br,
	}, nil
}

func (rs *remindService) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			rs.logger.Info("remind service closed")
			return nil
		case <-time.NewTicker(rs.freq).C:
			rs.remindSchedule(ctx)
		}
	}
}

func (rs *remindService) remindSchedule(ctx context.Context) {
	date := time.Now().Format("2006-01-02")
	events, err := rs.remind.GetEventsByDate(ctx, date)
	if err != nil {
		rs.logger.Warn(err)
		return
	}

	data, err := json.Marshal(&events)
	if err != nil {
		rs.logger.Warn(err)
		return
	}

	if err = rs.broker.Publish("remind", data); err != nil {
		rs.logger.Warn(err)
		return
	}

	rs.logger.WithFields("info", map[string]interface{}{
		"service": "remind",
		"bytes":   len(data),
	}, "successful published to broker")
}
