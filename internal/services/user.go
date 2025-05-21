package services

import (
	"errors"

	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/domain"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/dto"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/helper"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	hashPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	usr, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(usr.ID, usr.Email, usr.UserType)
}
func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exists with the provided user email")
	}
	// var token string
	if err := s.Auth.VerifyPassword(password, user.Password); err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}
func (s UserService) findUserByEmail(email string) (domain.User, error) {
	user, err := s.Repo.FindUser(email)
	return user, err
}
func (s UserService) GetVerificationCode(u domain.User) (int, error) {
	return 0, nil
}
func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}
func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}
func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}
func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}
func (s UserService) BecomeSeller(id uint, u domain.User) (string, error) {
	return "", nil
}
func (s UserService) FindCart(id uint) ([]any, error) {
	return nil, nil
}
func (s UserService) CreateCart(input any, u domain.User) ([]any, error) {
	return nil, nil
}
func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}
func (s UserService) GetOrders(u domain.User) ([]any, error) {
	return nil, nil
}
func (s UserService) GetOrderById(id uint, uId uint) (any, error) {
	return "", nil
}
