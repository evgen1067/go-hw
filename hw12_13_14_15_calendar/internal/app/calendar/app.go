package calendar

import (
	"context"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository/memory"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository/psql"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/server/grpcapi"
	httpApi "github.com/evgen1067/hw12_13_14_15_calendar/internal/server/httpapi"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/service"
	"os/signal"
	"syscall"
)

func Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var repo repository.EventsRepo
	conf := config.Configuration
	if conf.SQL {
		repo = psql.NewRepo()
	} else {
		repo = memory.NewRepo()
	}

	if r, ok := repo.(repository.DatabaseRepo); ok {
		err := r.Connect(ctx)
		if err != nil {
			return err
		}
		logger.Logger.Info("Database started.")
		defer r.Close()
	}

	services := service.NewServices(ctx, repo)

	httpSrv := httpApi.InitHTTP(services, conf)
	grpcSrv := grpcapi.InitGRPC(conf, services)

	errs := make(chan error)
	go func() {
		logger.Logger.Info("HTTP server started.")
		err := httpSrv.ListenAndServe()
		if err != nil {
			errs <- err
		}
	}()
	go func() {
		logger.Logger.Info("GRPC server started.")
		err := grpcSrv.ListenAndServe()
		if err != nil {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		logger.Logger.Error(err.Error())
		return err
	case <-ctx.Done():
		logger.Logger.Info("Shutdown with Signal.")
	}

	grpcSrv.Srv.GracefulStop()

	if err := httpSrv.Shutdown(ctx); err != nil {
		return err
	}

	logger.Logger.Info("Servers Exited Properly.")

	return nil
}
