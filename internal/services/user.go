package services

import (
	"errors"
	"time"

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

func (s UserService) IsVerifiedUser(id uint) bool {
	currentUser, err := s.Repo.FindUserById(id)
	return err == nil && currentUser.Verified
}

func (s UserService) GetVerificationCode(u domain.User) (int, error) {
	if s.IsVerifiedUser(u.ID) {
		return 0, errors.New("user already verified")
	}

	code, err := s.Auth.GenerateCode()

	if err != nil {
		return 0, err
	}

	user := domain.User{
		Expiry: time.Now().Add(time.Minute * 30),
		Code:   code,
	}

	if _, err := s.Repo.UpdateUser(u.ID, user); err != nil {
		return 0, err
	}

	return code, nil
}
func (s UserService) VerifyCode(id uint, code int) error {
	if s.IsVerifiedUser(id) {
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	userVerfied := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, userVerfied)
	if err != nil {
		return errors.New("unable to verify")
	}

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
func (s UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {
	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return "", err
	}

	if user.UserType == domain.SELLER {
		return "", errors.New("user is already seller")
	}

	seller, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.PhoneNumber,
		UserType:  domain.SELLER,
	})
	if err != nil {
		return "", err
	}

	token, err := s.Auth.GenerateToken(id, user.Email, seller.UserType)
	if err != nil {
		return "", err
	}

	if err := s.Repo.CreateBankAccount(domain.BankAccount{
		BankAccount: input.BankAccountNumber,
		SwiftCode:   input.SwiftCode,
		PaymentType: input.PaymentType,
		UserID:      id,
	}); err != nil {
		return "", errors.New("error while creating bank account " + err.Error())
	}

	return token, nil
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
