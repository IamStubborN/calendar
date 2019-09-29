package repository

import (
	"context"

	"github.com/IamStubborN/calendar/pkg/remind"

	"github.com/IamStubborN/calendar/models"
	"github.com/jmoiron/sqlx"
)

type remindRepository struct {
	pool *sqlx.DB
}

func NewRemindRepositoryPSQL(pool *sqlx.DB) remind.Repository {
	return &remindRepository{
		pool: pool,
	}
}

func (d *remindRepository) GetEventsByDate(ctx context.Context, date string) ([]*models.Event, error) {
	query := `select id, name, description, date from events where date=$1`

	rows, err := d.pool.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.Event
	for rows.Next() {
		var ev models.Event
		err = rows.Scan(&ev.ID, &ev.Name, &ev.Description, &ev.Date)

		events = append(events, &ev)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return events, nil
}
