package main

import (
	"log"

	"github.com/Pratam-Kalligudda/Ecommerce-go/config"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/api"
)

func main() {
	config, err := config.SetupEnv()
	if err != nil {
		log.Fatal("couldn't load env")
	}
	api.StartService(config)
}
