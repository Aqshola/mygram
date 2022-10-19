package repository

import (
	"errors"
	"mygram/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *entity.User) (*entity.User, error)
	// Update(user *entity.User) (*entity.User, error)
	// FindByEmail(email string) (*entity.User, error)
	// Delete(user *entity.User) error
	IsEmailExist(email string) bool
	IsUsernameExist(username string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Insert(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// func (r *userRepository) Update(user *entity.User) (*entity.User, error)

// func (r *userRepository) FindByEmail(email string) (*entity.User, error)

// func (r *userRepository) Delete(user *entity.User) error

func (r *userRepository) IsEmailExist(email string) bool {
	user := new(entity.User)
	err := r.db.Where("email = ?", email).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}

	}
	return true
}

func (r *userRepository) IsUsernameExist(username string) bool {
	user := new(entity.User)
	err := r.db.Where("username = ?", username).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}
