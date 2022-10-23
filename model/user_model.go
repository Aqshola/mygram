package model

import "time"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6" `
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Age      int    `json:"age" validate:"required,lte=100,gte=8"`
	Email    string `json:"email" validate:"required,email,max=191"`
	Password string `json:"password" validate:"required,min=6"`
	Username string `json:"username" validate:"required,max=10"`
}

type RegisterResponse struct {
	Age      int    `json:"age" `
	Email    string `json:"email"`
	Id       uint   `json:"id" `
	Username string `json:"username"`
}

type UpdateRequest struct {
	Email    string `json:"email" validate:"required,email,max=191"`
	Username string `json:"username" validate:"required,max=10"`
}

type UpdateResponse struct {
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	Id         uint      `json:"id"`
	Username   string    `json:"username"`
	Updated_at time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
