package repository

import (
	"errors"
	"mygram/entity"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Insert(comment *entity.Comment) (*entity.Comment, error)
	Update(comment *entity.Comment) (*entity.Comment, error)
	FindAll() (*[]entity.Comment, error)
	FindById(id uint) (*entity.Comment, error)
	Delete(id uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepostiory(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Insert(comment *entity.Comment) (*entity.Comment, error) {
	err := r.db.Table("comments").Create(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}
func (r *commentRepository) Update(comment *entity.Comment) (*entity.Comment, error) {
	err := r.db.Table("comments").Preload("users").Preload("photo").Where("id = ?", &comment.Id).Updates(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}
func (r *commentRepository) FindAll() (*[]entity.Comment, error) {
	var listComment []entity.Comment
	err := r.db.Table("comments").Preload("User").Preload("Photo").Find(&listComment).Error

	if err != nil {
		return &listComment, err
	}

	return &listComment, nil
}
func (r *commentRepository) FindById(id uint) (*entity.Comment, error) {
	var comment entity.Comment
	err := r.db.Table("comments").Where("id = ?", id).Take(&comment).Error
	if err != nil {
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &comment, errors.New("comment not found")
			}
			return &comment, errors.New("error while get comment")
		}
	}
	return &comment, nil
}
func (r *commentRepository) Delete(id uint) error {
	err := r.db.Table("comments").Where("id = ?", id).Delete(entity.Comment{}).Error

	if err != nil {
		return err
	}
	return nil
}
