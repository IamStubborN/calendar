package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/IamStubborN/calendar/api/models"
	"github.com/IamStubborN/calendar/api/pkg/event"
)

type EventsCache struct {
	sync.RWMutex
	storage map[uint64]*models.Event
}

func NewEventRepositoryCache() event.Repository {
	return &EventsCache{
		RWMutex: sync.RWMutex{},
		storage: make(map[uint64]*models.Event),
	}
}

func (e *EventsCache) Create(ctx context.Context, event *models.Event) (*models.Event, error) {
	e.Lock()
	defer e.Unlock()
	id := uint64(len(e.storage))

	event.ID = id
	e.storage[id] = event

	return event, nil
}

func (e *EventsCache) Read(ctx context.Context, eventID uint64) (*models.Event, error) {
	e.RLock()
	defer e.RUnlock()

	if uint64(len(e.storage)-1) < eventID {
		return nil, errors.New("event does not exists")
	}

	return e.storage[eventID], nil
}

func (e *EventsCache) Update(ctx context.Context, event *models.Event) (bool, error) {
	e.Lock()
	defer e.Unlock()

	if uint64(len(e.storage)-1) < event.ID {
		return false, errors.New("event does not exists")
	}

	e.storage[event.ID] = event

	return true, nil
}

func (e *EventsCache) Delete(ctx context.Context, eventID uint64) (bool, error) {
	e.Lock()
	defer e.Unlock()

	if len(e.storage) == 0 || uint64(len(e.storage)-1) < eventID {
		return false, errors.New("event does not exists")
	}

	delete(e.storage, eventID)

	return true, nil
}

func (e *EventsCache) GetEventsByDate(ctx context.Context, date string) ([]*models.Event, error) {
	var events []*models.Event

	for _, ev := range e.storage {
		if ev.Date == date {
			events = append(events, ev)
		}
	}

	return events, nil
}
