package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	t.Run("Test Config with valid path", func(t *testing.T) {
		Configuration, err := InitConfig("../../configs/config-test.json")
		require.NoError(t, err)

		require.Equal(t, "127.0.0.1", Configuration.DB.Host)
		require.Equal(t, "5432", Configuration.DB.Port)
		require.Equal(t, "test_user", Configuration.DB.User)
		require.Equal(t, "test_password", Configuration.DB.Password)
		require.Equal(t, "test_database", Configuration.DB.Database)

		require.Equal(t, Info, Configuration.Logger.Level)
		require.Equal(t, "out.log", Configuration.Logger.File)
	})

	t.Run("Test Config with invalid path", func(t *testing.T) {
		_, err := InitConfig("fail.json")
		require.Error(t, err)
	})
}
