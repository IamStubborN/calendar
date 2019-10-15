package event

import (
	"context"

	"github.com/IamStubborN/calendar/api/models"
)

type Repository interface {
	Create(ctx context.Context, ev *models.Event) (*models.Event, error)
	Read(ctx context.Context, eventID uint64) (*models.Event, error)
	Update(ctx context.Context, ev *models.Event) (bool, error)
	Delete(ctx context.Context, eventID uint64) (bool, error)
}
