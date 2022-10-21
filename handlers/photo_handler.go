package handlers

import (
	"mygram/helpers"
	"mygram/middlewares"
	"mygram/model"
	"mygram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type PhotoHandler struct {
	service service.PhotoService
}

func NewPhotoController(service service.PhotoService) *PhotoHandler {
	return &PhotoHandler{service: service}
}

func (controller *PhotoHandler) Route(route *gin.Engine) {
	photo := route.Group("/photos").Use(middlewares.Authentication())
	photo.POST("/", controller.AddPhoto)
	photo.GET("/", controller.GetAllPhoto)
	photo.PUT("/:photoId", controller.UpdatePhoto)
	photo.DELETE("/:photoId", controller.DeletePhoto)
}

func (controller *PhotoHandler) AddPhoto(ctx *gin.Context) {
	var addPhotoRequest model.AddPhotoRequest
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	ctx.ShouldBindJSON(&addPhotoRequest)

	valid, trans := helpers.Valid()
	err := valid.Struct(addPhotoRequest)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, e.Translate(trans), nil)
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	res, errAdd := controller.service.AddPhoto(userId, &addPhotoRequest)
	if errAdd != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errAdd.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusCreated, "Add new photo", res)
	ctx.JSON(http.StatusCreated, response)
}

func (controller *PhotoHandler) GetAllPhoto(ctx *gin.Context) {
	res, errGet := controller.service.GetAllPhoto()

	if errGet != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errGet.Error(), res)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success get all photo", res)
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller *PhotoHandler) UpdatePhoto(ctx *gin.Context) {
	ids := ctx.Param("photoId")

	idconvert, errconvert := strconv.Atoi(ids)
	if errconvert != nil {
		response := helpers.GenerateApiResponse(http.StatusBadGateway, "Unable to parse id", nil)
		ctx.JSON(http.StatusBadGateway, response)
	}

	var updatePhotoRequest model.UpdatePhotoRequest

	ctx.ShouldBindJSON(&updatePhotoRequest)

	valid, trans := helpers.Valid()
	err := valid.Struct(updatePhotoRequest)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, e.Translate(trans), nil)
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	res, errUpdate := controller.service.UpdatePhoto(uint(idconvert), &updatePhotoRequest)

	if errUpdate != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errUpdate.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Update Photo", res)
	ctx.JSON(http.StatusOK, response)

}

func (controller *PhotoHandler) DeletePhoto(ctx *gin.Context) {
	ids := ctx.Param("photoId")

	parsedId, errParsed := strconv.Atoi(ids)

	if errParsed != nil {
		response := helpers.GenerateApiResponse(http.StatusBadGateway, "Invalid id", nil)
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	res, errDelete := controller.service.DeletePhoto(uint(parsedId))
	if errDelete != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errDelete.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Photto has been deleted", res)
	ctx.JSON(http.StatusOK, response)

}