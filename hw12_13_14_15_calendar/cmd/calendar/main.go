package main

import (
	"flag"
	"log"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/app/calendar"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/local.json", "Path to json configuration file")
}

func main() {
	flag.Parse()
	err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Error during logger initialization: %s", err)
		return
	}
	_, err = config.InitConfig(configFile)
	if err != nil {
		logger.Logger.Error("Error when reading the configuration file: " + err.Error())
		return
	}
	app := calendar.InitApp()
	err = app.Start()
	if err != nil {
		logger.Logger.Error("Error when launching the application: " + err.Error())
		return
	}
}
