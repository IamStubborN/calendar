package remind

import (
	"context"

	"github.com/IamStubborN/calendar/models"
)

type Repository interface {
	GetEventsByDate(ctx context.Context, date string) ([]*models.Event, error)
}
