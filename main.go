package main

import (
	"fmt"
	"log"
	"mygram/config"
	"mygram/handlers"
	"mygram/repository"
	"mygram/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Unable to load env")
	}

	idb := config.StartDB()
	router := gin.Default()

	//INIT USER
	userRepo := repository.NewUserRepository(idb)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserController(userService)

	//INIT PHOTO
	photoRepo := repository.NewPhotoRepository(idb)
	photoService := service.NewPhotoService(photoRepo)
	photoHandler := handlers.NewPhotoController(photoService)

	//ROUTE
	userHandler.Route(router)
	photoHandler.Route(router)

	router.Run(":8080")
	fmt.Println("Server Running")
}