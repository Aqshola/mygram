package repository

import (
	"errors"
	"mygram/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Insert(photo *entity.Photo) (*entity.Photo, error)
	Update(Photo *entity.Photo) (*entity.Photo, error)
	FindAll() (*[]entity.Photo, error)
	Delete(id uint) error
	FindById(id uint) (*entity.Photo, error)
}

type repository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindById(id uint) (*entity.Photo, error) {
	var photo entity.Photo
	err := r.db.Table("photos").Where("id = ?", id).Take(&photo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &photo, errors.New("photo not found")
		}
		return &photo, errors.New("error while get photo")
	}

	return &photo, err
}

func (r *repository) Insert(photo *entity.Photo) (*entity.Photo, error) {
	err := r.db.Table("photos").Create(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) Update(photo *entity.Photo) (*entity.Photo, error) {
	err := r.db.Table("photos").Preload("users").Where("id = ?", &photo.Id).Updates(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) FindAll() (*[]entity.Photo, error) {
	var listPhoto []entity.Photo
	err := r.db.Table("photos").Preload("User").Find(&listPhoto).Error

	if err != nil {
		return &listPhoto, err
	}

	return &listPhoto, nil
}

func (r *repository) Delete(id uint) error {
	err := r.db.Table("photos").Where("id = ?", id).Error

	if err != nil {
		return err
	}
	return nil
}
