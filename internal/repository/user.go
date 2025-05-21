package repository

import (
	"errors"
	"log"

	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) CreateUser(u domain.User) (domain.User, error) {
	if err := r.db.Create(&u).Error; err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return u, nil
}
func (r userRepository) FindUser(email string) (user domain.User, err error) {
	if err := r.db.Find(&user, "email=?", email).Error; err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("couldnt find user")
	}
	return user, nil
}
func (r userRepository) FindUserById(id uint) (user domain.User, err error) {
	if err = r.db.Find(&user, "id=?", id).Error; err != nil {
		log.Printf("find user by id error %v", err)
		return domain.User{}, errors.New("couldnt find user by id")
	}
	return
}
func (r userRepository) UpdateUser(id uint, u domain.User) (user domain.User, err error) {
	if err = r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error; err != nil {
		log.Printf("couldnt update user error %v", err)
		return domain.User{}, errors.New("couldnt update user")
	}
	return
}
