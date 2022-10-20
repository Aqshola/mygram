package middlewares

import (
	"mygram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.ValidateJwt(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err":     "Unauthenicated",
				"message": err.Error(),
			})
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
