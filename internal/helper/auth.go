package helper

import (
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/domain"
	"github.com/gofiber/fiber"
)

type Auth struct {
	Secret string
}

func SetupAuth(secret string) Auth {
	return Auth{Secret: secret}
}

func (a Auth) CreateHashPassword(pass string) (string, error) {
	return "", nil
}

func (a Auth) VerifyPassword(pass, hashPass string) error {
	return nil
}

func (a Auth) GenerateToken(u domain.User) (string, error) {
	return "", nil
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	return nil
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) error {
	return nil
}

func (a Auth) GenerateCode() (int, error) {
	return -1, nil
}
