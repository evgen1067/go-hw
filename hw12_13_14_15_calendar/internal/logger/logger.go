package logger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = &zap.Logger{}

func InitLogger() error {
	var (
		defaultOutputPaths = []string{"out.log"}
		err                error
	)
	var level zap.AtomicLevel
	switch config.Configuration.Logger.Level {
	case config.Error:
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case config.Warn:
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case config.Info:
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case config.Debug:
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	var file []string
	if config.Configuration.Logger.File == "" {
		file = defaultOutputPaths
	} else {
		file = append(file, config.Configuration.Logger.File)
	}
	pathDir := "logs"
	if _, err := os.Stat(pathDir); os.IsNotExist(err) {
		err := os.Mkdir(pathDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	for i := range file {
		file[i] = filepath.Join(pathDir, file[i])
	}
	file = append(file, "stdout")

	cfg := zap.Config{
		Level:       level,
		Encoding:    "console",
		OutputPaths: file,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			TimeKey:     "time",
			EncodeLevel: CustomEncodeLevel,
			EncodeTime:  CustomEncodeTime,
		},
	}

	Logger, err = cfg.Build()
	if err != nil {
		return err
	}

	return nil
}

func CustomEncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + l.String() + "]")
}

func CustomEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2/Jan/2006:15:04:05 -0700"))
}
