package service

import (
	"errors"
	"mygram/dto"
	"mygram/entity"
	"mygram/helpers"
	"mygram/repository"
	"time"
)

type SocialService interface {
	CreateSocial(userId uint, request *dto.CreateSocialRequest) (dto.CreateSocialResponse, error)
	GetAllSocial(userId uint) ([]dto.GetAllSocialResponse, error)
	UpdateSocial(id uint, request *dto.UpdateSocialRequest) (dto.UpdateSocialResponse, error)
	DeleteSocial(id uint) (dto.DeleteSocialResponse, error)
}

type socialService struct {
	repository     repository.SocialRepository
	userRepository repository.UserRepository
}

func NewSocialService(repository repository.SocialRepository, userRepository repository.UserRepository) *socialService {
	return &socialService{repository: repository, userRepository: userRepository}
}

func (r *socialService) CreateSocial(userId uint, request *dto.CreateSocialRequest) (dto.CreateSocialResponse, error) {
	var social entity.SocialMedia = entity.SocialMedia{
		Name:             request.Name,
		Social_Media_Url: request.Social_media_url,
		User_Id:          userId,
		Created_at:       time.Now(),
	}

	parsedUrl, errUrl := helpers.ParseAndValidateUrl(request.Social_media_url)

	social.Social_Media_Url = parsedUrl

	if errUrl != nil {
		return dto.CreateSocialResponse{}, errors.New("invalid social url")
	}
	res, err := r.repository.Insert(&social)
	if err != nil {
		return dto.CreateSocialResponse{}, err
	}

	return dto.CreateSocialResponse{
		Id:               res.Id,
		Name:             res.Name,
		Social_media_url: res.Social_Media_Url,
		User_id:          res.User_Id,
		Created_at:       res.Created_at,
	}, nil
}
func (r *socialService) GetAllSocial(userId uint) ([]dto.GetAllSocialResponse, error) {
	var listSocial []dto.GetAllSocialResponse = []dto.GetAllSocialResponse{}

	res, err := r.repository.FindAllByUser(userId)
	if err != nil {
		return listSocial, err
	}

	for _, v := range *res {
		listSocial = append(listSocial, dto.GetAllSocialResponse{
			Id:               v.Id,
			Name:             v.Name,
			Social_media_url: v.Social_Media_Url,
			User_id:          v.User_Id,
			Created_at:       v.Created_at,
			User: dto.UserResponse{
				Id:       v.User.Id,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		})

	}
	return listSocial, nil
}
func (r *socialService) UpdateSocial(id uint, request *dto.UpdateSocialRequest) (dto.UpdateSocialResponse, error) {
	social, errGet := r.repository.FindById(id)

	if errGet != nil {
		return dto.UpdateSocialResponse{}, errGet
	}

	parsedUrl, errUrl := helpers.ParseAndValidateUrl(request.Social_media_url)

	if errUrl != nil {
		return dto.UpdateSocialResponse{}, errors.New("Invalid social url")
	}

	social.Name = request.Name
	social.Social_Media_Url = parsedUrl

	res, errUpdate := r.repository.Update(social)

	if errUpdate != nil {
		return dto.UpdateSocialResponse{}, errUpdate
	}

	return dto.UpdateSocialResponse{
		Id:               res.Id,
		Name:             res.Name,
		Social_media_url: res.Social_Media_Url,
		User_id:          res.User_Id,
		Updated_at:       res.Updated_at,
	}, nil

}
func (r *socialService) DeleteSocial(id uint) (dto.DeleteSocialResponse, error) {
	errDelete := r.repository.Delete(id)

	if errDelete != nil {
		return dto.DeleteSocialResponse{
			Message: "Error while delete social",
		}, errDelete
	}

	return dto.DeleteSocialResponse{
		Message: "Your social media has been successfully deleted",
	}, nil
}
