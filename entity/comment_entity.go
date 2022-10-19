package entity

import "time"

type Comment struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	User_Id    uint      `json:"user_id" `
	Photo_Id   uint      `json:"photo_id" `
	Message    string    `json:"message" gorm:"type varchar(191)"`
	Created_at time.Time `json:"created_at" gorm:"type datetime"`
	Updated_at time.Time `json:"updated_at" gorm:"type datetime"`
	User       *User
	Photo      *Photo
}
