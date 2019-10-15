package repository

import (
	"context"
	"errors"

	"github.com/IamStubborN/calendar/api/models"
	"github.com/IamStubborN/calendar/api/pkg/event"
	"github.com/jmoiron/sqlx"
)

type eventRepository struct {
	pool *sqlx.DB
}

func NewEventRepositoryPSQL(pool *sqlx.DB) event.Repository {
	return &eventRepository{
		pool: pool,
	}
}

func (er *eventRepository) Create(ctx context.Context, ev *models.Event) (*models.Event, error) {
	query := `
	insert into events("name", description, "date") 
		values (:name, :description, :date) returning id
	`

	argsQ := map[string]interface{}{
		"name":        ev.Name,
		"description": ev.Description,
		"date":        ev.Date,
	}

	stmt, err := er.pool.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var id uint64
	err = stmt.QueryRowxContext(ctx, argsQ).Scan(&id)
	if err != nil {
		return nil, err
	}

	ev.ID = id

	return ev, nil
}

func (er *eventRepository) Read(ctx context.Context, eventID uint64) (*models.Event, error) {
	query := `select id, "name", description, "date" from events where id=$1`

	var ev models.Event

	err := er.pool.QueryRowxContext(ctx, query, eventID).Scan(&ev.ID, &ev.Name, &ev.Description, &ev.Date)
	if err != nil {
		return nil, err
	}

	return &ev, nil
}

func (er *eventRepository) Update(ctx context.Context, ev *models.Event) (bool, error) {
	query := `update events set "name"=:name, description=:description, "date"=:date 
		where id=:id`

	argsQ := map[string]interface{}{
		"id":          ev.ID,
		"name":        ev.Name,
		"description": ev.Description,
		"date":        ev.Date,
	}

	res, err := er.pool.NamedExecContext(ctx, query, argsQ)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, errors.New("can't update event")
	}

	return true, nil
}

func (er *eventRepository) Delete(ctx context.Context, eventID uint64) (bool, error) {
	query := `delete from events where id=$1`

	res, err := er.pool.ExecContext(ctx, query, eventID)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, errors.New("can't delete event")
	}

	return true, nil
}

func (er *eventRepository) GetEventsByDate(ctx context.Context, date string) ([]*models.Event, error) {
	query := `select id, name, description, date from events where date=$1`

	rows, err := er.pool.Query(query, date)
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