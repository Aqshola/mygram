package service

import (
	"errors"
	"mygram/entity"
	"mygram/helpers"
	"mygram/model"
	"mygram/repository"
	"time"
)

type UserService interface {
	Register(request *model.RegisterRequest) (model.RegisterResponse, error)
	// Login(request *model.LoginRequest) (model.LoginResponse, error)
	// UpdateUser(request *model.UpdateRequest) (model.UpdateResponse, error)
	// DeleteUser(id string) (interface{}, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Register(request *model.RegisterRequest) (model.RegisterResponse, error) {
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
		return model.RegisterResponse{
			Email: request.Email,
		}, errors.New("email already used")
	}

	if isUsernameExist {
		return model.RegisterResponse{
			Username: request.Username,
		}, errors.New("username already used")
	}

	newUser, errorInsert := s.repository.Insert(&user)
	if errorInsert != nil {
		return model.RegisterResponse{
			Age:      request.Age,
			Email:    request.Email,
			Username: request.Username,
		}, errorInsert
	}

	return model.RegisterResponse{
		Email:    newUser.Email,
		Age:      newUser.Age,
		Id:       newUser.Id,
		Username: newUser.Username,
	}, nil
}

// func (s *userService) Login(request *model.LoginRequest) (model.LoginResponse, error)
// func (s *userService) UpdateUser(request *model.UpdateRequest) (model.UpdateResponse, error)
// func (s *userService) DeleteUser(id string) (interface{}, error)
