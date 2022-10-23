package service

import (
	"errors"
	"mygram/dto"
	"mygram/entity"
	"mygram/helpers"
	"mygram/repository"
	"time"
)

type UserService interface {
	Register(request *dto.RegisterRequest) (dto.RegisterResponse, error)
	Login(request *dto.LoginRequest) (dto.LoginResponse, error)
	UpdateUser(id uint, request *dto.UpdateRequest) (dto.UpdateResponse, error)
	DeleteUser(id uint) (dto.DeleteUserResponse, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Register(request *dto.RegisterRequest) (dto.RegisterResponse, error) {
	user := entity.User{
		Age:        request.Age,
		Email:      request.Email,
		Username:   request.Username,
		Password:   helpers.HashPass(request.Password),
		Created_at: time.Now(),
	}

	isEmailExist := s.repository.IsEmailExist(user.Email)
	isUsernameExist := s.repository.IsUsernameExist(user.Username)

	if isEmailExist {
		return dto.RegisterResponse{
			Email: request.Email,
		}, errors.New("email already used")
	}

	if isUsernameExist {
		return dto.RegisterResponse{
			Username: request.Username,
		}, errors.New("username already used")
	}

	newUser, errorInsert := s.repository.Insert(&user)
	if errorInsert != nil {
		return dto.RegisterResponse{
			Age:      request.Age,
			Email:    request.Email,
			Username: request.Username,
		}, errorInsert
	}

	return dto.RegisterResponse{
		Email:    newUser.Email,
		Age:      newUser.Age,
		Id:       newUser.Id,
		Username: newUser.Username,
	}, nil
}

func (s *userService) Login(request *dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.repository.FindByEmail(request.Email)

	if err != nil {
		return dto.LoginResponse{
			Token: "",
		}, err
	}

	comparePass := helpers.ComparePass(request.Password, user.Password)
	if !comparePass {
		return dto.LoginResponse{}, errors.New("wrong password")
	}

	_token := helpers.GenerateJWT(user.Id, user.Email)

	return dto.LoginResponse{
		Token: _token,
	}, nil

}

func (s *userService) UpdateUser(id uint, request *dto.UpdateRequest) (dto.UpdateResponse, error) {
	user, err := s.repository.FindById(id)

	if err != nil {
		return dto.UpdateResponse{}, err
	}

	user.Email = request.Email
	user.Username = request.Username
	user.Updated_at = time.Now()

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return dto.UpdateResponse{}, err
	}

	return dto.UpdateResponse{
		Id:         updateUser.Id,
		Email:      updateUser.Email,
		Username:   updateUser.Username,
		Age:        updateUser.Age,
		Updated_at: updateUser.Updated_at,
	}, nil
}

func (s *userService) DeleteUser(id uint) (dto.DeleteUserResponse, error) {
	_, errUser := s.repository.FindById(id)
	if errUser != nil {
		return dto.DeleteUserResponse{
			Message: "Error while delete",
		}, errUser
	}
	errDelete := s.repository.Delete(id)

	if errDelete != nil {
		return dto.DeleteUserResponse{
			Message: "Error while delete",
		}, errDelete
	}

	return dto.DeleteUserResponse{
		Message: "Your account has been successfully deleted",
	}, nil

}
