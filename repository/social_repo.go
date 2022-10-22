package repository

import (
	"errors"
	"mygram/entity"

	"gorm.io/gorm"
)

type SocialRepository interface {
	Insert(social *entity.SocialMedia) (*entity.SocialMedia, error)
	Update(social *entity.SocialMedia) (*entity.SocialMedia, error)
	FindAll() (*[]entity.SocialMedia, error)
	FindById(id uint) (*entity.SocialMedia, error)
	Delete(id uint) error
}

type socialRepository struct {
	db *gorm.DB
}

func NewSocialRepository(db *gorm.DB) *socialRepository {
	return &socialRepository{db}
}

func (r *socialRepository) Insert(social *entity.SocialMedia) (*entity.SocialMedia, error) {

	err := r.db.Table("social_media").Create(&social).Error

	if err != nil {
		return social, err
	}

	return social, nil
}

func (r *socialRepository) Update(social *entity.SocialMedia) (*entity.SocialMedia, error) {
	err := r.db.Table("social_media").Where("id = ?", social.Id).Updates(&social).Error
	if err != nil {
		return social, err
	}

	return social, nil
}

func (r *socialRepository) FindAll() (*[]entity.SocialMedia, error) {
	var listSocial []entity.SocialMedia
	err := r.db.Table("social_media").Preload("User").Find(&listSocial).Error

	if err != nil {
		return &listSocial, err
	}

	return &listSocial, nil
}

func (r *socialRepository) FindById(id uint) (*entity.SocialMedia, error) {
	var social entity.SocialMedia

	err := r.db.Table("social_media").Where("id = ?", id).Take(&social).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &social, errors.New("comment not found")
		}
		return &social, errors.New("error while get comment")
	}
	return &social, nil
}

func (r *socialRepository) Delete(id uint) error {
	err := r.db.Table("social_media").Where("id = ?", id).Delete(entity.SocialMedia{}).Error

	if err != nil {
		return err
	}

	return nil
}
