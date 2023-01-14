package logger

import (
	"bytes"
	"os"
	"testing"

	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	var err error
	var lines []byte
	fileName := "logs/out.log"
	t.Run("test logger without file name", func(t *testing.T) {
		messages := []string{
			"Test logger error message",
			"Test logger warn message",
			"Test logger info message",
		}
		err = InitLogger()
		require.NoError(t, err)

		l := Logger
		l.Error(messages[0])
		l.Warn(messages[1])
		l.Info(messages[2])

		lines, err = os.ReadFile(fileName)
		require.NoError(t, err)

		err = os.RemoveAll("logs")
		require.NoError(t, err)

		require.Contains(t, string(bytes.Split(lines, []byte("\n"))[0]), messages[0])
		require.Contains(t, string(bytes.Split(lines, []byte("\n"))[1]), messages[1])
		require.Contains(t, string(bytes.Split(lines, []byte("\n"))[2]), messages[2])

		require.NoError(t, err)
	})

	t.Run("test logger with file name", func(t *testing.T) {
		fileName := "data.log"
		config.Configuration.Logger.File = fileName
		messages := []string{
			"Test logger error message",
			"Test logger warn message",
			"Test logger info message",
		}
		err = InitLogger()
		require.NoError(t, err)

		l := Logger
		l.Error(messages[0])
		l.Warn(messages[1])
		l.Info(messages[2])

		lines, err = os.ReadFile("logs/" + fileName)
		require.NoError(t, err)

		err = os.RemoveAll("logs")

		require.Contains(t, string(bytes.Split(lines, []byte("\n"))[0]), messages[0])
		require.Contains(t, string(bytes.Split(lines, []byte("\n"))[1]), messages[1])
		require.Contains(t, string(bytes.Split(lines, []byte("\n"))[2]), messages[2])

		require.NoError(t, err)
	})
}
