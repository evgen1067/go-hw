package main

import (
	"context"
	"flag"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/rabbit/producer"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository/psql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/local.json", "Path to json configuration file")
}

func main() {
	flag.Parse()
	err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Error during logger initialization: %s", err.Error())
		return
	}
	cfg, err := config.InitConfig(configFile)
	if err != nil {
		logger.Logger.Error("Error when reading the configuration file: " + err.Error())
		return
	}
	logger.Logger.Info("The scheduler has started working")
	defer logger.Logger.Info("The scheduler has finished its work")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := psql.NewRepo()

	repo, ok := repo.(repository.DatabaseRepo)
	if !ok {
		return
	}
	err = repo.Connect(ctx)
	if err != nil {
		logger.Logger.Error("Error when connecting to the database: " + err.Error())
		return
	}
	logger.Logger.Info("Database started.")
	defer repo.Close()

	prod := producer.NewProducer(cfg.AMQP.Uri, cfg.AMQP.Queue)
	err = prod.Start()
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	defer prod.Stop()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go work(ctx, repo, prod)

	<-done
}

func work(ctx context.Context, repo repository.DatabaseRepo, prod *producer.Producer) {
	ticker := time.NewTicker(1 * time.Minute)
	for {
		err := repo.ClearOldEvents(ctx) // точно работает
		logger.Logger.Info("test")
		if err != nil {
			logger.Logger.Error("Error when clearing old events: " + err.Error())
		}
		notices, err := repo.SchedulerList(ctx) // это тоже по идее должна работать
		if err != nil {
			logger.Logger.Error("Error when receiving notifications from the database: " + err.Error())
		}
		for _, v := range notices {
			n, err := v.MarshalJSON()
			if err != nil {
				logger.Logger.Error("Error when marshaling notifications: " + err.Error())
			}
			err = prod.Publish(ctx, n)
			if err != nil {
				logger.Logger.Error("Error when publishing a notification by the producer: " + err.Error())
			}
		}
		select {
		case <-ticker.C:
			continue
		case <-ctx.Done():
			return
		}
	}
}
