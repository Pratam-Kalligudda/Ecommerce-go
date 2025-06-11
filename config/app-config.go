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
	godotenv.Load()

	serverPort := os.Getenv("HOST")
	if len(serverPort) < 1 {
		return AppConfig{}, errors.New("couldnt load env host")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("couldnt load env dsn")
	}

	secret := os.Getenv("APP_SECRET")
	if len(secret) < 1 {
		return AppConfig{}, errors.New("couldnt load env secret")
	}

	return AppConfig{
		ServerPort: serverPort,
		Dsn:        dsn,
		AppSecret:  secret,
	}, nil
}
