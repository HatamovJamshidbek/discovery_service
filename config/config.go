package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	CompositionMetaDataPort string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	LOG_PATH string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.CompositionMetaDataPort = cast.ToString(coalesce("COMPOSITION_META_DATA_PORT", ":8084"))

	config.PostgresHost = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(coalesce("DB_PORT", 5433))
	config.PostgresUser = cast.ToString(coalesce("DB_USER", "postgres"))
	config.PostgresPassword = cast.ToString(coalesce("DB_PASSWORD", "1111"))
	config.PostgresDatabase = cast.ToString(coalesce("DB_NAME", "composition_meta_datas"))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
