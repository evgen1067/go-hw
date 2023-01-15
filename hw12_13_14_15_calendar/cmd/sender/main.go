package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/rabbit/consumer"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configFile string
	messages   <-chan amqp091.Delivery
)

func init() {
	flag.StringVar(&configFile, "config", "configs/local.json", "Path to json configuration file")
}

func main() {
	flag.Parse()
	err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Error during logger initialization: %s", err)
	}
	cfg, err := config.InitConfig(configFile)
	if err != nil {
		logger.Logger.Error("Error when reading the configuration file: " + err.Error())
	}
	cons := consumer.NewConsumer(cfg.AMQP.Uri, cfg.AMQP.Queue)
	err = cons.Start()
	if err != nil {
		logger.Logger.Error(err.Error())
	}
	defer cons.Stop()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	messages, err = cons.Consume()
	if err != nil {
		logger.Logger.Error(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go work(ctx)

	<-done
}

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-messages:
			var notice repository.Notice
			err := notice.UnmarshalJSON(msg.Body)
			if err != nil {
				logger.Logger.Error("Error when unmarshaling notifications: " + err.Error())
				continue
			}
			logger.Logger.Info(fmt.Sprintf("ID: %v, Title: %v, Datetime: %v, OwnerID: %v",
				notice.EventID, notice.Title, notice.Datetime, notice.OwnerID))
		}
	}
}
