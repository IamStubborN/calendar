package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IamStubborN/calendar/api/pkg/logger"
	"github.com/IamStubborN/calendar/api/worker"
)

type App struct {
	Logger  logger.UseCase
	Workers []worker.Worker
}

func NewApp() *App {
	var app App
	cfg := initializeConfig()
	app.Logger = initializeLogger(cfg)
	pool, err := initializeSQLConn(cfg, app.Logger)
	if err != nil {
		app.Logger.Fatal(err)
	}

	er := initializeEventRepository(cfg, pool)
	app.Workers = append(app.Workers, initializeEventService(app.Logger, er))

	return &app
}

func (app *App) Run() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	for _, w := range app.Workers {
		wg.Add(1)

		go func(w worker.Worker) {
			defer wg.Done()
			if err := w.Run(ctx); err != nil {
				app.Logger.Warn(err)
			}
		}(w)
	}

	gracefulShutdown(cancel)
	wg.Wait()
}

func gracefulShutdown(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	close(c)
	cancel()
}
