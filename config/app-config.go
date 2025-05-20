package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (config AppConfig, err error) {
	// if os.Getenv("APP_ENV") == "Dev" {
	godotenv.Load()
	// }

	serverPort := os.Getenv("HTTP_PORT")

	if len(serverPort) < 1 {
		return AppConfig{}, errors.New("couldnt load http port from env")
	}

	return AppConfig{ServerPort: serverPort}, nil
}
