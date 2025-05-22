package helper

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		secret: s,
	}
}

func (a Auth) CreateHashedPassword(pass string) (string, error) {
	if len(pass) < 6 {
		return "", errors.New("password lenght should be at least 6 characters long")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", errors.New("password hash failed")
	}

	return string(hashPass), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("required inputs are missing to generate token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "ecommerce-go",
		"sub":   id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", errors.New("signing token failed")
	}

	return tokenStr, nil
}

func (a Auth) VerifyPassword(pass, hashPass string) error {
	if len(pass) < 6 {
		return errors.New("password lenght should be at least 6 characters long")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass)); err != nil {
		return errors.New("password does not match")
	}
	return nil
}

func (a Auth) VerifyToken(token string) (domain.User, error) {
	tokenArr := strings.Split(token, " ")
	log.Println("token : " + token)
	if len(tokenArr) != 2 {
		log.Println("len of token not 2")
		return domain.User{}, nil
	}

	tokenStr := tokenArr[1]

	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	t, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", t.Header)
		}
		return []byte(a.secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{}
		user.ID = uint(claims["sub"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)
		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]
	if len(authHeader) == 0 {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  "null",
		})
	}
	user, err := a.VerifyToken(authHeader[0])
	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	}
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)
}

func (a Auth) GenerateCode() (int, error) {
	return RandomNumbers(6)
}
