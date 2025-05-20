package api

import (
	"github.com/Pratam-Kalligudda/Ecommerce-go/config"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/api/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

func StartService(config config.AppConfig) {
	app := fiber.New()
	rh := &rest.RestHandler{
		App: app,
	}
	SetupRoutes(rh)
	app.Listen(config.ServerPort)
}

func SetupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
	//catalog routes
	//transactions routes
}
