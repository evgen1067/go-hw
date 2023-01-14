package config

import (
	"os"
)

type LoggerLevel string

const (
	Error LoggerLevel = "error"
	Warn  LoggerLevel = "warn"
	Info  LoggerLevel = "info"
	Debug LoggerLevel = "debug"
)

type Config struct {
	Logger struct {
		Level LoggerLevel `json:"level"`
		File  string      `json:"file"`
	} `json:"logger"`
	HTTP struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"http"`
	GRPC struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"grpc"`
	SQL bool `json:"sql"`
	DB  struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"db"`
}

var Configuration = &Config{}

func InitConfig(filePath string) (*Config, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = Configuration.UnmarshalJSON(bytes)
	if err != nil {
		return nil, err
	}
	return Configuration, nil
}
