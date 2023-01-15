package calendar

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository/memory"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository/psql"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/server/grpcapi"
	httpApi "github.com/evgen1067/hw12_13_14_15_calendar/internal/server/httpapi"
)

type App struct {
	ctx    context.Context
	config *config.Config
	repo   repository.EventsRepo
	http   *httpApi.Server
	grpc   *grpcapi.Server
}

func InitApp() *App {
	var repo repository.EventsRepo
	conf := config.Configuration
	if conf.SQL {
		repo = psql.NewRepo()
	} else {
		repo = memory.NewRepo()
	}
	ctx := context.Background()
	httpSrv := httpApi.InitHTTP(ctx, repo, conf)
	grpcSrv := grpcapi.InitGRPC(ctx, repo, conf)
	return &App{
		ctx:    ctx,
		config: conf,
		repo:   repo,
		http:   httpSrv,
		grpc:   grpcSrv,
	}
}

func (app *App) Start() error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if repo, ok := app.repo.(repository.DatabaseRepo); ok {
		err := repo.Connect(app.ctx)
		if err != nil {
			return err
		}
		logger.Logger.Info("Database started.")
		defer repo.Close()
	}

	errs := make(chan error)
	go func() {
		logger.Logger.Info("HTTP server started.")
		err := app.http.Srv.ListenAndServe()
		if err != nil {
			errs <- err
		}
	}()
	go func() {
		logger.Logger.Info("GRPC server started.")
		err := app.grpc.ListenAndServe()
		if err != nil {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		logger.Logger.Error(err.Error())
		return err
	case <-done:
		logger.Logger.Info("Shutdown with Signal.")
	}

	ctx, cancel := context.WithTimeout(app.ctx, 3*time.Second)
	defer cancel()

	app.grpc.Srv.GracefulStop()

	// now close the server gracefully ("shutdown")
	if err := app.http.Srv.Shutdown(ctx); err != nil {
		return err
	}

	logger.Logger.Info("Servers Exited Properly.")

	return nil
}
