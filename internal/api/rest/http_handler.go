package rest

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Handler struct {
	App *fiber.App
	DB  *gorm.DB
}
