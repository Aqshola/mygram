package dto

import "time"

type AddPhotoRequest struct {
	Title     string `json:"title" validate:"required,max=191"`
	Caption   string `json:"caption" validate:"required,max=191"`
	Photo_url string `json:"photo_url" validate:"required,max=191"`
}

type AddPhotoResponse struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    uint      `json:"user_id"`
	Created_at time.Time `json:"created_at"`
}

type GetAllPhotoResponse struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    uint      `json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User       UserResponse
}

type UpdatePhotoRequest struct {
	Title     string `json:"title" validate:"required,max=191"`
	Caption   string `json:"caption" validate:"required,max=191"`
	Photo_url string `json:"photo_url" validate:"required,max=191"`
}

type UpdatePhotoResponse struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    uint      `json:"user_id"`
	Updated_at time.Time `json:"updated_at"`
}

type DeletePhotoResponse struct {
	Message string `json:"message"`
}

type PhotoResponse struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_id   uint   `json:"user_id"`
}
