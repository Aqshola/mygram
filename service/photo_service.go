package service

import (
	"errors"
	"mygram/dto"
	"mygram/entity"
	"mygram/helpers"
	"mygram/repository"
	"path/filepath"
	"time"
)

type PhotoService interface {
	AddPhoto(userId uint, request *dto.AddPhotoRequest) (dto.AddPhotoResponse, error)
	GetAllPhoto() ([]dto.GetAllPhotoResponse, error)
	UpdatePhoto(id uint, request *dto.UpdatePhotoRequest) (dto.UpdatePhotoResponse, error)
	DeletePhoto(id uint) (dto.DeletePhotoResponse, error)
}

type photoService struct {
	repository     repository.PhotoRepository
	userRepository repository.UserRepository
}

func NewPhotoService(repository repository.PhotoRepository, userRepository repository.UserRepository) *photoService {
	return &photoService{repository: repository, userRepository: userRepository}
}

func (s *photoService) AddPhoto(userId uint, request *dto.AddPhotoRequest) (dto.AddPhotoResponse, error) {

	var photo entity.Photo = entity.Photo{
		Title:      request.Title,
		Photo_url:  request.Photo_url,
		Caption:    request.Caption,
		Created_at: time.Now(),
		User_Id:    userId,
	}

	_, errCheckUser := s.userRepository.FindById(photo.User_Id)

	if errCheckUser != nil {
		return dto.AddPhotoResponse{
			Title:     request.Title,
			Caption:   request.Caption,
			Photo_url: request.Photo_url,
		}, errCheckUser
	}

	fileExt := filepath.Ext(photo.Photo_url)

	parsedUrl, errUrl := helpers.ParseAndValidateUrl(request.Photo_url)
	photo.Photo_url = parsedUrl

	if errUrl != nil && (fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg") {
		return dto.AddPhotoResponse{}, errors.New("invalid photo url")
	}

	res, err := s.repository.Insert(&photo)
	if err != nil {
		return dto.AddPhotoResponse{}, errors.New("error while add photo")
	}

	return dto.AddPhotoResponse{
		Id:         res.Id,
		Title:      res.Title,
		Caption:    res.Caption,
		Photo_url:  res.Photo_url,
		User_id:    res.User_Id,
		Created_at: res.Created_at,
	}, nil
}

func (s *photoService) GetAllPhoto() ([]dto.GetAllPhotoResponse, error) {
	var listPhoto []dto.GetAllPhotoResponse = []dto.GetAllPhotoResponse{}

	res, errGet := s.repository.FindAll()
	if errGet != nil {
		return listPhoto, errors.New("error while get photo")
	}

	for _, v := range *res {
		listPhoto = append(listPhoto, dto.GetAllPhotoResponse{
			Id:         v.Id,
			Title:      v.Title,
			Caption:    v.Caption,
			Photo_url:  v.Photo_url,
			User_id:    v.User_Id,
			Created_at: v.Created_at,
			Updated_at: v.Updated_at,
			User: dto.UserResponse{
				Id:       v.User.Id,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		})

	}

	return listPhoto, nil
}

func (s *photoService) UpdatePhoto(id uint, request *dto.UpdatePhotoRequest) (dto.UpdatePhotoResponse, error) {
	photo, errGet := s.repository.FindById(id)

	fileExt := filepath.Ext(request.Photo_url)

	parsedUrl, errUrl := helpers.ParseAndValidateUrl(request.Photo_url)

	if errUrl != nil && (fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg") {
		return dto.UpdatePhotoResponse{}, errors.New("invalid photo url")
	}

	if errGet != nil {
		return dto.UpdatePhotoResponse{}, errGet
	}

	photo.Title = request.Title
	photo.Caption = request.Caption
	photo.Photo_url = parsedUrl
	photo.Updated_at = time.Now()

	photoUpdate, errUpdate := s.repository.Update(photo)
	if errUpdate != nil {
		return dto.UpdatePhotoResponse{}, errUpdate
	}

	return dto.UpdatePhotoResponse{
		Id:         photo.Id,
		Title:      photoUpdate.Title,
		Caption:    photoUpdate.Caption,
		Photo_url:  photoUpdate.Photo_url,
		User_id:    photoUpdate.User_Id,
		Updated_at: photoUpdate.Updated_at,
	}, nil

}

func (s *photoService) DeletePhoto(id uint) (dto.DeletePhotoResponse, error) {

	errDelete := s.repository.Delete(id)

	if errDelete != nil {
		return dto.DeletePhotoResponse{
			Message: "Fail to delete photo",
		}, errDelete
	}
	return dto.DeletePhotoResponse{
		Message: "Your photo has been successfully deleted!",
	}, nil
}
