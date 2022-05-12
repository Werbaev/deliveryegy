package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	ServiceHost string
	ServicePort string

	PostgresHost     string
	PostgresPort     int
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string

	LogLevel string

	// context timeout in seconds
	CtxTimeout int
}

// Load loads environment vars and inflates Config
func Load() Config {
	config := Config{}

	config.ServiceHost = cast.ToString(getOrReturnDefault("SERVICE_HOST", "0.0.0.0"))
	config.ServicePort = cast.ToString(getOrReturnDefault("SERVICE_PORT", ":5001"))

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DB", "delivery"))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "postgres"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
