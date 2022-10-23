package main

import (
	"fmt"
	"log"
	"mygram/config"
	_ "mygram/docs"
	"mygram/handlers"
	"mygram/repository"
	"mygram/service"
	"os"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           MyGram Example API
// @version         1.0
// @host      localhost:8080
// @BasePath  /

// @securitydefinitions.apikey Authorization
// @in header
// @name Authorization

func main() {

	PORT := "8080"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load production env")
	}

	if os.Getenv("APP_ENV") == "production" {
		PORT = os.Getenv("PORT")
	}

	idb := config.StartDB()
	router := gin.Default()

	//INIT USER
	userRepo := repository.NewUserRepository(idb)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserController(userService)

	//INIT PHOTO
	photoRepo := repository.NewPhotoRepository(idb)
	photoService := service.NewPhotoService(photoRepo, userRepo)
	photoHandler := handlers.NewPhotoController(photoService)

	//INIT COMMENT
	commentRepo := repository.NewCommentRepostiory(idb)
	commentService := service.NewCommentService(commentRepo, userRepo, photoRepo)
	commentHandler := handlers.NewCommentController(commentService)

	//INIT SOCIAL
	socialRepo := repository.NewSocialRepository(idb)
	socialService := service.NewSocialService(socialRepo, userRepo)
	socialHandler := handlers.NewSocialHandler(socialService)

	//ROUTE
	userHandler.Route(router)
	photoHandler.Route(router)
	commentHandler.Route(router)
	socialHandler.Route(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":" + PORT)
	fmt.Println("Server Running")
}
