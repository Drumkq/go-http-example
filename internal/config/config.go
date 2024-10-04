package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Debug             bool
	CurrentApiVersion string
	DbHost            string
	DbPort            string
	DbUser            string
	DbPassword        string
	DbName            string
	DbSSLMode         string
}

func Init() (Config, error) {
	var cfg Config

	isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return Config{}, fmt.Errorf("enviroment vars loading error: %w", err)
	}

	envFile := ".env"
	if isDebug {
		envFile = ".env.development"
	}

	err = godotenv.Load(envFile)
	if err != nil {
		return Config{}, fmt.Errorf("enviroment vars loading error: %w", err)
	}

	// Initialization of the config
	cfg.CurrentApiVersion = os.Getenv("CURRENT_API_VERSION")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbPort = os.Getenv("DB_PORT")
	cfg.DbUser = os.Getenv("DB_USER")
	cfg.DbPassword = os.Getenv("DB_PASSWORD")
	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbSSLMode = os.Getenv("DB_SSLMODE")

	return cfg, nil
}
