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
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open((postgres.Open(psqlInfo)), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connect database", err)
	}

	fmt.Println("Database connected")
	if os.Getenv("APP_ENV") == "production" {
		db.AutoMigrate(entity.User{}, entity.Photo{}, entity.SocialMedia{}, entity.Comment{})
	} else {
		db.Debug().AutoMigrate(entity.User{}, entity.Photo{}, entity.SocialMedia{}, entity.Comment{})
	}

	return db
}

func CallDB() *gorm.DB {
	return db
}
