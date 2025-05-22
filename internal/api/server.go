package api

import (
	"log"

	"github.com/Pratam-Kalligudda/Ecommerce-go/config"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/api/rest/handlers"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/domain"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartService(config config.AppConfig) {
	app := fiber.New()
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error : %v\n", err)
	}

	log.Println("database connected")

	//db automigrate
	err = db.AutoMigrate(&domain.User{}, &domain.BankAccount{})
	if err != nil {
		log.Fatalf("error on running migration : %v", err.Error())
	}
	log.Println("migration was succesfully")

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: auth,
	}

	SetupRoutes(rh)

	app.Listen(config.ServerPort)
}

func SetupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
	//catalog routes
	//transactions routes
}
