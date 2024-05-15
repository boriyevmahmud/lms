package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDatabase string
	ServiceName      string

	RedisHost     string
	RedisPort     string
	RedisPassword string
}

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file, err: ", err)
	}
	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "lms"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "makhmud"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1"))
	cfg.ServiceName = "LMS"

	cfg.RedisHost = "localhost" //cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	cfg.RedisPort = "6379"      //cast.ToString(getOrReturnDefault("REDIS_PORT", "6379"))
	cfg.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "password"))
	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {

	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
