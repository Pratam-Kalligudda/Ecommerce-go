package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	AppSecret  string
}

func SetupEnv() (config AppConfig, err error) {
	// if os.Getenv("APP_ENV") == "Dev" {
	godotenv.Load()
	// }

	serverPort := os.Getenv("HTTP_PORT")
	if len(serverPort) < 1 {
		return AppConfig{}, errors.New("couldnt load http port from env")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("couldnt load dsn from env")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("couldnt load app secret from env")
	}

	return AppConfig{ServerPort: serverPort, Dsn: dsn, AppSecret: appSecret}, nil
}
