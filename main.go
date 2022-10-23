package main

import (
	"fmt"
	"log"
	"mygram/config"
	_ "mygram/docs"
	"mygram/handlers"
	"mygram/middlewares"
	"mygram/repository"
	"mygram/service"
	"net/http"
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
	if os.Getenv("APP_ENV") == "production" {
		PORT = os.Getenv("PORT")
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Unable to load production env")
		}
	}

	idb := config.StartDB()
	router := gin.Default()

	router.Use(middlewares.CORS())
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
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
	userHandler.Route(router)
	photoHandler.Route(router)
	commentHandler.Route(router)
	socialHandler.Route(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":" + PORT)
	fmt.Println("Server Running")
}
