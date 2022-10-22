package service

import (
	"mygram/entity"
	"mygram/model"
	"mygram/repository"
	"time"
)

type CommentService interface {
	CreateComment(userId uint, request *model.CreateCommentRequest) (model.CreateCommentResponse, error)
	UpdateComment(id uint, request *model.UpdateCommentRequest) (model.UpdateCommentResponse, error)
	GetAllComment() ([]model.GetAllCommentResponse, error)
	DeleteComment(id uint) (model.DeleteCommentResponse, error)
}

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) *commentService {
	return &commentService{repository: repository}
}

func (s *commentService) CreateComment(userId uint, request *model.CreateCommentRequest) (model.CreateCommentResponse, error) {
	var comment entity.Comment = entity.Comment{
		Message:    request.Message,
		Photo_Id:   request.Photo_id,
		User_Id:    userId,
		Created_at: time.Now(),
	}

	res, errCreate := s.repository.Insert(&comment)

	if errCreate != nil {
		return model.CreateCommentResponse{
			Message:  request.Message,
			Photo_id: request.Photo_id,
		}, errCreate
	}

	return model.CreateCommentResponse{
		Id:         res.Id,
		Message:    res.Message,
		Photo_id:   res.Photo_Id,
		Created_at: res.Created_at,
	}, nil

}
func (s *commentService) UpdateComment(id uint, request *model.UpdateCommentRequest) (model.UpdateCommentResponse, error) {
	comment, errDetail := s.repository.FindById(id)
	comment.Message = request.Message

	if errDetail != nil {
		return model.UpdateCommentResponse{
			Message: request.Message,
		}, errDetail
	}

	res, errUpdate := s.repository.Update(comment)

	if errUpdate != nil {
		return model.UpdateCommentResponse{
			Message: request.Message,
		}, errUpdate
	}

	return model.UpdateCommentResponse{
		Id:         res.Id,
		Photo_Id:   res.Photo_Id,
		Message:    res.Message,
		User_id:    res.User_Id,
		Updated_at: time.Now(),
	}, nil

}
func (s *commentService) GetAllComment() ([]model.GetAllCommentResponse, error) {
	var listComment []model.GetAllCommentResponse = []model.GetAllCommentResponse{}

	res, errGet := s.repository.FindAll()

	if errGet != nil {
		return listComment, errGet
	}

	for _, v := range *res {
		listComment = append(listComment, model.GetAllCommentResponse{
			Id:         v.Id,
			Message:    v.Message,
			Photo_id:   v.Photo_Id,
			User_id:    v.User_Id,
			Created_at: v.Created_at,
			Updated_at: v.Updated_at,
			User: &model.UserResponse{
				Id:       v.User.Id,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
			Photo: &model.PhotoResponse{
				Id:        v.Photo.Id,
				Title:     v.Photo.Title,
				Caption:   v.Photo.Caption,
				Photo_url: v.Photo.Photo_url,
				User_id:   v.User_Id,
			},
		})
	}
	return listComment, nil

}
func (s *commentService) DeleteComment(id uint) (model.DeleteCommentResponse, error) {
	_, errDetail := s.repository.FindById(id)

	if errDetail != nil {
		return model.DeleteCommentResponse{
			Message: "Error while delete",
		}, errDetail
	}
	err := s.repository.Delete(id)

	if err != nil {
		return model.DeleteCommentResponse{
			Message: "Error while delete",
		}, errDetail
	}

	return model.DeleteCommentResponse{
		Message: "Your comment has been successfully deleted",
	}, nil
}
