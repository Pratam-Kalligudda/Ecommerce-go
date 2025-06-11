package api

import (
	"log"

	"github.com/Pratam-Kalligudda/Ecommerce-go/config"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/domain"
	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error : %v", err.Error())
	}
	log.Printf("database connected")

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("error while migrating %v", err.Error())
	}
	log.Println("migration was succesful")

	app.Listen(config.ServerPort)
}
