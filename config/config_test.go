package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var filename = "./../.env"

func TestLoadConfig_Fail(t *testing.T) {
	// os.Setenv("POSTGRES_HOST", "localhost")
	filename := "./.env"
	err := LoadConfig(filename)
	require.NotNil(t, err)
}

func TestLoadConfig_Success(t *testing.T) {

	err := LoadConfig(filename)
	require.Nil(t, err)
}

func TestGetString_Default(t *testing.T) {
	dbHost := GetString(POSTGRES_HOST, "default")
	require.Equal(t, "default", dbHost)
}

func TestGetString_Success(t *testing.T) {
	os.Setenv("POSTGRES_HOST", "localhost")
	err := LoadConfig(filename)

	require.Nil(t, err)

	dbHost := GetString(POSTGRES_HOST, "default")
	require.Equal(t, "localhost", dbHost)

}

func TestGetInt8_Default(t *testing.T) {
	maxConn := GetInt8(POSTGRES_MAX_OPEN_CONN, 25)
	require.Equal(t, int8(25), maxConn)
}

func TestGetInt8_Success(t *testing.T) {
	os.Setenv(string(POSTGRES_MAX_OPEN_CONN), "10")
	maxConn := GetInt8(POSTGRES_MAX_OPEN_CONN, 25)
	require.Equal(t, int8(10), maxConn)
}
