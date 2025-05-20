package services

import "github.com/Pratam-Kalligudda/Ecommerce-go/internal/domain"

type UserService struct {
}

func (s UserService) Signup(input any) (string, error) {
	return "", nil
}
func (s UserService) Login(input any) (string, error) {
	return "", nil
}
func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	return nil, nil
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
