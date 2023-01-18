package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	t.Run("Test Config with valid path", func(t *testing.T) {
		Configuration, err := InitConfig("../../configs/config-test.json")
		require.NoError(t, err)

		require.Equal(t, "0.0.0.0", Configuration.DB.Host)
		require.Equal(t, "6000", Configuration.DB.Port)
		require.Equal(t, "go_user", Configuration.DB.User)
		require.Equal(t, "go_password", Configuration.DB.Password)
		require.Equal(t, "events_db", Configuration.DB.Database)

		require.Equal(t, Info, Configuration.Logger.Level)
		require.Equal(t, "out.log", Configuration.Logger.File)
	})

	t.Run("Test Config with invalid path", func(t *testing.T) {
		_, err := InitConfig("fail.json")
		require.Error(t, err)
	})
}
