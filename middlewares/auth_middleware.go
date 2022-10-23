package middlewares

import (
	"errors"
	"mygram/config"
	"mygram/helpers"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.ValidateJwt(ctx)

		if err != nil {
			response := helpers.GenerateApiResponse(http.StatusUnauthorized, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload := verifyToken.(jwt.MapClaims)
		expiredAt := payload["expiredAt"]
		userId := payload["id"]

		if expiredAt == nil {
			response := helpers.GenerateApiResponse(http.StatusUnauthorized, "Invalid Token", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if time.Now().Unix() > int64(expiredAt.(float64)) {
			response := helpers.GenerateApiResponse(http.StatusUnauthorized, "Token Expired", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tempCheck struct {
			Id uint `json:"id"`
		}

		errUser := config.CallDB().Table("users").Select("id").Where("id = ?", uint(userId.(float64))).Take(&tempCheck).Error

		if errUser != nil {
			if errors.Is(errUser, gorm.ErrRecordNotFound) {
				response := helpers.GenerateApiResponse(http.StatusUnauthorized, "User Invalid", nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			response := helpers.GenerateApiResponse(http.StatusInternalServerError, "Internal Server Error", nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}

func Authorization(service string, idparam string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))

		ids := ctx.Param(idparam)
		idconvert, errconvert := strconv.Atoi(ids)
		if errconvert != nil {
			response := helpers.GenerateApiResponse(http.StatusBadGateway, "Unable to parse id", nil)
			ctx.AbortWithStatusJSON(http.StatusBadGateway, response)
			return
		}

		var serviceUserId struct {
			User_id uint `json:"user_id"`
		}

		errService := config.CallDB().Table(service).Select("user_id").Where("id = ?", uint(idconvert)).Take(&serviceUserId).Error

		if errService != nil && errors.Is(errService, gorm.ErrRecordNotFound) {
			response := helpers.GenerateApiResponse(http.StatusNotFound, "Content not found", nil)
			ctx.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}

		if serviceUserId.User_id != userId {
			response := helpers.GenerateApiResponse(http.StatusForbidden, "Forbidden access", nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		ctx.Next()
	}
}
