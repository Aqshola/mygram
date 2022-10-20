package repository

import (
	"errors"
	"mygram/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id uint) (*entity.User, error)
	Insert(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Delete(id uint) error
	IsEmailExist(email string) bool
	IsUsernameExist(username string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindById(id uint) (*entity.User, error) {
	var user entity.User

	err := r.db.Table("users").Where("id = ?", id).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, errors.New("user not found")
		}
		return &user, errors.New("unable to delete")
	}

	return &user, nil
}

func (r *userRepository) Insert(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user *entity.User) (*entity.User, error) {
	err := r.db.Table("users").Where("id = ?", &user.Id).Updates(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.Table("users").Where("email = ?", email).Take(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, errors.New("user not found")
		}
	}
	return &user, nil
}

func (r *userRepository) Delete(id uint) error {
	err := r.db.Table("users").Where("id = ?", id).Delete(entity.User{}).Error

	if err != nil {
		return errors.New("error while delete")
	}

	return nil

}

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
