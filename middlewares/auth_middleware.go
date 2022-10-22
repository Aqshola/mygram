package middlewares

import (
	"mygram/helpers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

		expiredAt := verifyToken.(jwt.MapClaims)["expiredAt"]

		if expiredAt == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err":     "Unauthenicated",
				"message": "Invalid Token",
			})
		}

		if time.Now().Unix() > int64(expiredAt.(float64)) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err":     "Unauthenicated",
				"message": "Token Expired",
			})
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
