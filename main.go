package main

import (
	"log"

	"github.com/Pratam-Kalligudda/Ecommerce-go/config"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/api"
)

func main() {
	config, err := config.SetupEnv()
	if err != nil {
		log.Fatal("couldnt setup env : " + err.Error())
	}
	api.StartServer(config)
}
