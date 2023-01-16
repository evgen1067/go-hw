package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/rabbit/producer"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository/psql"
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
	}
	cfg, err := config.InitConfig(configFile)
	if err != nil {
		logger.Logger.Error("Error when reading the configuration file: " + err.Error())
	}
	logger.Logger.Info("The scheduler has started working")
	defer logger.Logger.Info("The scheduler has finished its work")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := psql.NewRepo()

	err = repo.Connect(ctx)
	if err != nil {
		logger.Logger.Error("Error when connecting to the database: " + err.Error())
	}
	logger.Logger.Info("Database started.")
	defer repo.Close()

	prod := producer.NewProducer(cfg.AMQP.URI, cfg.AMQP.Queue)
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
	ticker := time.NewTicker(5 * time.Minute)
	for {
		err := repo.ClearOldEvents(ctx)
		if err != nil {
			logger.Logger.Error("Error when clearing old events: " + err.Error())
		}
		notices, err := repo.SchedulerList(ctx)
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
