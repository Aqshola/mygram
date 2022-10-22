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
	FindById(id uint) (*entity.Photo, error)
	Delete(id uint) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) FindById(id uint) (*entity.Photo, error) {
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

func (r *photoRepository) Insert(photo *entity.Photo) (*entity.Photo, error) {
	err := r.db.Table("photos").Create(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *photoRepository) Update(photo *entity.Photo) (*entity.Photo, error) {
	err := r.db.Table("photos").Preload("users").Where("id = ?", &photo.Id).Updates(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *photoRepository) FindAll() (*[]entity.Photo, error) {
	var listPhoto []entity.Photo
	err := r.db.Table("photos").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,email,username")
	}).Find(&listPhoto).Error

	if err != nil {
		return &listPhoto, err
	}

	return &listPhoto, nil
}

func (r *photoRepository) Delete(id uint) error {
	err := r.db.Table("photos").Where("id = ?", id).Delete(entity.Photo{}).Error

	if err != nil {
		return err
	}
	return nil
}
