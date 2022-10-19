package entity

import "time"

type User struct {
	Id           uint          `gorm:"primaryKey" json:"id"`
	Username     string        `json:"username" gorm:"type varchar(10);unique;not null"`
	Email        string        `json:"email" gorm:"type varchar(191);not null;unique"`
	Password     string        `json:"password"  gorm:"type varchar(191); not null"`
	Age          int           `json:"age"`
	Created_at   time.Time     `json:"created_at"`
	Updated_at   time.Time     `json:"updated_at"`
	Photos       []Photo       `gorm:"foreignKey User_Id; references Id;constraint:onDelete:CASCADE" json:"photos"`
	Social_Media []SocialMedia `gorm:"foreignKey User_Id; references Id;constraint:onDelete:CASCADE" json:"social_media"`
	Comments     []Comment     `gorm:"foreignKey User_Id; references Id;constraint:onDelete:CASCADE" json:"comments"`
}
