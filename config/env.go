package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPath string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBPath: getEnv("DB_PATH", "db/app.db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
