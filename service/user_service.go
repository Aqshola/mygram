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
	Login(request *model.LoginRequest) (model.LoginResponse, error)
	UpdateUser(id uint, request *model.UpdateRequest) (model.UpdateResponse, error)
	DeleteUser(id uint) (model.DeleteUserResponse, error)
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

func (s *userService) Login(request *model.LoginRequest) (model.LoginResponse, error) {
	user, err := s.repository.FindByEmail(request.Email)

	if err != nil {
		return model.LoginResponse{
			Token: "",
		}, err
	}

	comparePass := helpers.ComparePass(request.Password, user.Password)
	if !comparePass {
		return model.LoginResponse{}, errors.New("wrong password")
	}

	_token := helpers.GenerateJWT(user.Id, user.Email)

	return model.LoginResponse{
		Token: _token,
	}, nil

}

func (s *userService) UpdateUser(id uint, request *model.UpdateRequest) (model.UpdateResponse, error) {
	user, err := s.repository.FindById(id)

	if err != nil {
		return model.UpdateResponse{}, err
	}

	user.Email = request.Email
	user.Username = request.Username
	user.Updated_at = time.Now()

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return model.UpdateResponse{}, err
	}

	return model.UpdateResponse{
		Id:         updateUser.Id,
		Email:      updateUser.Email,
		Username:   updateUser.Username,
		Age:        updateUser.Age,
		Updated_at: updateUser.Updated_at,
	}, nil
}

func (s *userService) DeleteUser(id uint) (model.DeleteUserResponse, error) {
	_, errUser := s.repository.FindById(id)
	if errUser != nil {
		return model.DeleteUserResponse{
			Message: "Error while delete",
		}, errUser
	}
	errDelete := s.repository.Delete(id)

	if errDelete != nil {
		return model.DeleteUserResponse{
			Message: "Error while delete",
		}, errDelete
	}

	return model.DeleteUserResponse{
		Message: "Your account has been successfully deleted",
	}, nil

}
