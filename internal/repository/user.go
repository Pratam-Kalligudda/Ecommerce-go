package repository

import (
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
	UpdateUser(u domain.User) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r userRepository) CreateUser(u domain.User) (domain.User, error) {
	return domain.User{}, nil
}
func (r userRepository) FindUserByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}
func (r userRepository) FindUserById(id uint) (domain.User, error) {
	return domain.User{}, nil
}
func (r userRepository) UpdateUser(u domain.User) (domain.User, error) {
	return domain.User{}, nil
}
