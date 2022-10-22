package model

import "time"

type CreateCommentRequest struct {
	Message  string `json:"message" validate:"required,max=191"`
	Photo_id uint   `json:"photo_id" validate:"required"`
}

type CreateCommentResponse struct {
	Id         uint      `json:"id"`
	Message    string    `json:"message"`
	Photo_id   uint      `json:"photo_id"`
	Created_at time.Time `json:"created_at"`
}

type GetAllCommentResponse struct {
	Id         uint      `json:"id"`
	Message    string    `json:"message"`
	Photo_id   uint      `json:"photo_id"`
	User_id    uint      `json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User       *UserResponse
	Photo      *PhotoResponse
}

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"required,max=191"`
}

type UpdateCommentResponse struct {
	Id         uint      `json:"id"`
	Photo_Id   uint      `json:"photo_id"`
	Message    string    `json:"message"`
	User_id    uint      `json:"user_id"`
	Updated_at time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
