package handlers

import (
	"mygram/helpers"
	"mygram/middlewares"
	"mygram/model"
	"mygram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type SocialHandler struct {
	service service.SocialService
}

func NewSocialHandler(service service.SocialService) *SocialHandler {
	return &SocialHandler{service: service}
}

func (controller *SocialHandler) Route(route *gin.Engine) {
	social := route.Group("/socialmedias").Use(middlewares.Authentication())
	social.POST("/", controller.CreateSocial)
	social.GET("/", controller.GetAllSocial)
	social.Use(middlewares.Authorization("social_media", "socialMediaId")).PUT("/:socialMediaId", controller.UpdateSocial)
	social.Use(middlewares.Authorization("social_media", "socialMediaId")).DELETE("/:socialMediaId", controller.DeleteSocial)
}

func (controller *SocialHandler) CreateSocial(ctx *gin.Context) {
	var createSocialRequest model.CreateSocialRequest
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	ctx.ShouldBindJSON(&createSocialRequest)
	errValid := helpers.CheckValid(createSocialRequest)
	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res, errCreate := controller.service.CreateSocial(userId, &createSocialRequest)

	if errCreate != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errCreate.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success create social", model.CreateSocialResponse{
		Id:               res.Id,
		Name:             res.Name,
		Social_media_url: res.Social_media_url,
		User_id:          res.User_id,
		Created_at:       res.Created_at,
	})
	ctx.JSON(http.StatusOK, response)
}
func (controller *SocialHandler) GetAllSocial(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	res, err := controller.service.GetAllSocial(userId)

	if err != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, err.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helpers.GenerateApiResponse(http.StatusOK, "Success get all comment", res)
	ctx.JSON(http.StatusOK, response)
}

func (controller *SocialHandler) UpdateSocial(ctx *gin.Context) {
	ids := ctx.Param("socialMediaId")
	idconvert, errconvert := strconv.Atoi(ids)
	if errconvert != nil {
		response := helpers.GenerateApiResponse(http.StatusBadGateway, "Unable to parse id", nil)
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	var updateSocialRequest model.UpdateSocialRequest
	ctx.ShouldBindJSON(&updateSocialRequest)

	errValid := helpers.CheckValid(updateSocialRequest)
	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res, errUpdate := controller.service.UpdateSocial(uint(idconvert), &updateSocialRequest)

	if errUpdate != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errUpdate.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success update social", res)
	ctx.JSON(http.StatusOK, response)

}
func (controller *SocialHandler) DeleteSocial(ctx *gin.Context) {
	ids := ctx.Param("socialMediaId")
	idconvert, errconvert := strconv.Atoi(ids)
	if errconvert != nil {
		response := helpers.GenerateApiResponse(http.StatusBadGateway, "Unable to parse id", nil)
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	res, errDelete := controller.service.DeleteSocial(uint(idconvert))
	if errDelete != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errDelete.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success delete data", res)
	ctx.JSON(http.StatusOK, response)
}
