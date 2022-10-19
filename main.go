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

	//INIT USER
	idb := config.StartDB()
	router := gin.Default()
	userRepo := repository.NewUserRepository(idb)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserController(userService)

	userHandler.Route(router)

	router.Run(":8080")
	fmt.Println("Server Running")
}
