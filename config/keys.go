package config

type ConfigKey string

const (
	POSTGRES_HOST          ConfigKey = "POSTGRES_HOST"
	POSTGRES_PORT          ConfigKey = "POSTGRES_PORT"
	POSTGRES_USER          ConfigKey = "POSTGRES_USER"
	POSTGRES_PASS          ConfigKey = "POSTGRES_PASS"
	POSTGRES_DB_NAME       ConfigKey = "POSTGRES_DB_NAME"
	POSTGRES_SSLMODE       ConfigKey = "POSTGRES_SSLMODE"
	POSTGRES_MAX_OPEN_CONN ConfigKey = "POSTGRES_MAX_OPEN_CONN"
	POSTGRES_MAX_LIFETIME  ConfigKey = "POSTGRES_MAX_LIFETIME"
	POSTGRES_MAX_IDLE_CONN ConfigKey = "POSTGRES_MAX_IDLE_CONN"
	POSTGRES_MAX_IDLE_TIME ConfigKey = "POSTGRES_MAX_IDLE_TIME"
)
