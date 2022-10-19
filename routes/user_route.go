package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(route *gin.Engine, idb *gorm.DB) {
	user := route.Group("/users")
	user.POST("/register")
	user.POST("/login")
	user.Use().PUT("/:userId")
	user.Use().DELETE("/:userId")
}
