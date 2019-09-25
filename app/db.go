package app

import (
	"errors"
	"time"

	"github.com/IamStubborN/calendar/pkg/logger"

	"github.com/IamStubborN/calendar/config"
	"github.com/jmoiron/sqlx"
)

func initializeSQLConn(cfg *config.Config, logger logger.Repository) (*sqlx.DB, error) {
	pool, err := sqlx.Open("postgres", cfg.Storage.DSN)
	if err != nil {
		return nil, err
	}

	if err := retryConnect(pool, cfg.Storage.Retry, logger); err != nil {
		return nil, err
	}

	return pool, nil
}

func retryConnect(pool *sqlx.DB, fatalRetry int, logger logger.Repository) error {
	var retryCount int
	for range time.NewTicker(time.Second).C {
		if fatalRetry == retryCount {
			return errors.New("can't connect to database")
		}

		retryCount++
		if err := pool.Ping(); err != nil {
			logger.WithFields("info", map[string]interface{}{
				"status": "retrying",
				"try":    retryCount,
			}, "connect to db")

			continue
		}

		logger.WithFields("info", map[string]interface{}{
			"status": "connected",
		}, "connect to db")
		break
	}

	return nil
}
