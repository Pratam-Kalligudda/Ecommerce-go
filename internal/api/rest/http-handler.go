package rest

import (
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App  *fiber.App
	DB   *gorm.DB
	Auth helper.Auth
}
