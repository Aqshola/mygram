package config

import (
	"fmt"
	"log"

	"mygram/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func StartDB() *gorm.DB {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open((postgres.Open(psqlInfo)), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connect database", err)
	}

	fmt.Println("Database connected")
	db.Debug().AutoMigrate(entity.User{}, entity.Photo{}, entity.SocialMedia{}, entity.Comment{})

	return db
}
