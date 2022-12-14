package service

import (
	"mygram/dto"
	"mygram/entity"
	"mygram/repository"
	"time"
)

type CommentService interface {
	CreateComment(userId uint, request *dto.CreateCommentRequest) (dto.CreateCommentResponse, error)
	GetAllComment() ([]dto.GetAllCommentResponse, error)
	UpdateComment(id uint, request *dto.UpdateCommentRequest) (dto.UpdateCommentResponse, error)
	DeleteComment(id uint) (dto.DeleteCommentResponse, error)
}

type commentService struct {
	repository      repository.CommentRepository
	photoRepository repository.PhotoRepository
	userRepository  repository.UserRepository
}

func NewCommentService(repository repository.CommentRepository, userRepository repository.UserRepository, photoRepository repository.PhotoRepository) *commentService {
	return &commentService{repository: repository, userRepository: userRepository, photoRepository: photoRepository}
}

func (s *commentService) CreateComment(userId uint, request *dto.CreateCommentRequest) (dto.CreateCommentResponse, error) {

	_, errCheck := s.photoRepository.FindById(request.Photo_id)

	if errCheck != nil {
		return dto.CreateCommentResponse{}, errCheck
	}
	var comment entity.Comment = entity.Comment{
		Message:    request.Message,
		Photo_Id:   request.Photo_id,
		User_Id:    userId,
		Created_at: time.Now(),
	}

	res, errCreate := s.repository.Insert(&comment)

	if errCreate != nil {
		return dto.CreateCommentResponse{
			Message:  request.Message,
			Photo_id: request.Photo_id,
		}, errCreate
	}

	return dto.CreateCommentResponse{
		Id:         res.Id,
		Message:    res.Message,
		Photo_id:   res.Photo_Id,
		Created_at: res.Created_at,
	}, nil

}
func (s *commentService) UpdateComment(id uint, request *dto.UpdateCommentRequest) (dto.UpdateCommentResponse, error) {

	comment, errDetail := s.repository.FindById(id)
	if errDetail != nil {
		return dto.UpdateCommentResponse{
			Message: request.Message,
		}, errDetail
	}

	comment.Message = request.Message

	res, errUpdate := s.repository.Update(comment)
	if errUpdate != nil {
		return dto.UpdateCommentResponse{
			Message: request.Message,
		}, errUpdate
	}

	return dto.UpdateCommentResponse{
		Id:         res.Id,
		Photo_Id:   res.Photo_Id,
		Message:    res.Message,
		User_id:    res.User_Id,
		Updated_at: time.Now(),
	}, nil

}
func (s *commentService) GetAllComment() ([]dto.GetAllCommentResponse, error) {
	var listComment []dto.GetAllCommentResponse = []dto.GetAllCommentResponse{}

	res, errGet := s.repository.FindAll()

	if errGet != nil {
		return listComment, errGet
	}

	for _, v := range *res {
		listComment = append(listComment, dto.GetAllCommentResponse{
			Id:         v.Id,
			Message:    v.Message,
			Photo_id:   v.Photo_Id,
			User_id:    v.User_Id,
			Created_at: v.Created_at,
			Updated_at: v.Updated_at,
			User: &dto.UserResponse{
				Id:       v.User.Id,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
			Photo: &dto.PhotoResponse{
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
func (s *commentService) DeleteComment(id uint) (dto.DeleteCommentResponse, error) {

	_, errDetail := s.repository.FindById(id)

	if errDetail != nil {
		return dto.DeleteCommentResponse{
			Message: "Error while delete",
		}, errDetail
	}

	err := s.repository.Delete(id)

	if err != nil {
		return dto.DeleteCommentResponse{
			Message: "Error while delete",
		}, errDetail
	}

	return dto.DeleteCommentResponse{
		Message: "Your comment has been successfully deleted",
	}, nil
}
