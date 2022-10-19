package entity

import "time"

type SocialMedia struct {
	Id               uint      `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name" gorm:"type varchar(50);not null"`
	Social_Media_Url string    `json:"social_media_url" gorm:"type varchar(191); not null"`
	User_Id          uint      `json:"user_id"`
	Created_at       time.Time `json:"created_at" gorm:"type datetime"`
	Updated_at       time.Time `json:"updated_at" gorm:"type datetime"`
	User             *User
}
