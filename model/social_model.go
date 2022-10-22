package model

import "time"

type CreateSocialRequest struct {
	Name             string `json:"name" validate:"required,max=191"`
	Social_media_url string `json:"social_media_url" validate:"required,max=191"`
}

type CreateSocialResponse struct {
	Id               uint      `json:"id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	User_id          uint      `json:"user_id"`
	Created_at       time.Time `json:"created_at"`
}
type GetAllSocialResponse struct {
	Id               uint         `json:"id"`
	Name             string       `json:"name"`
	Social_media_url string       `json:"social_media_url"`
	User_id          uint         `json:"user_id"`
	Created_at       time.Time    `json:"created_at"`
	User             UserResponse `json:"user"`
}

type UpdateSocialRequest struct {
	Name             string `json:"name" validate:"required,max=191"`
	Social_media_url string `json:"social_media_url" validate:"required,max=191"`
}

type UpdateSocialResponse struct {
	Id               uint      `json:"id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	User_id          uint      `json:"user_id"`
	Updated_at       time.Time `json:"updated_at"`
}

type DeleteSocialResponse struct {
	Message string `json:"message"`
}
