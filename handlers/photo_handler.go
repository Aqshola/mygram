package handlers

import (
	"mygram/dto"
	"mygram/helpers"
	"mygram/middlewares"
	"mygram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	photo.Use(middlewares.Authorization("photos", "photoId")).PUT("/:photoId", controller.UpdatePhoto)
	photo.Use(middlewares.Authorization("photos", "photoId")).DELETE("/:photoId", controller.DeletePhoto)
}

// @Summary      Add new photo
// @Description Add new photo
// @Tags         Photo
// @Security Authorization
// @Accept json
// @Produce json
// @Param addPhotoRequest body dto.AddPhotoRequest true "Add photo body"
// @Success 201 {object}  helpers.ApiResponse{data=dto.AddPhotoResponse} "Success"
// @Failure 422 {object}  helpers.ApiResponse
// @Router /photos/ [post]
func (controller *PhotoHandler) AddPhoto(ctx *gin.Context) {
	var addPhotoRequest dto.AddPhotoRequest
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	ctx.ShouldBindJSON(&addPhotoRequest)

	errValid := helpers.CheckValid(addPhotoRequest)
	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
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

// @Summary      Get All Photo
// @Description Get All photo
// @Tags         Photo
// @Security Authorization
// @Produce json
// @Success 200 {object}  helpers.ApiResponse{data=[]dto.GetAllPhotoResponse} "Success"
// @Failure 422 {object}  helpers.ApiResponse
// @Router /photos/ [get]
func (controller *PhotoHandler) GetAllPhoto(ctx *gin.Context) {
	res, errGet := controller.service.GetAllPhoto()

	if errGet != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errGet.Error(), res)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success get all photo", res)
	ctx.JSON(http.StatusOK, response)

}

// @Summary      Update Photo
// @Description Update Photo
// @Tags         Photo
// @Security Authorization
// @Accept json
// @Produce json
// @Param updatePhotoRequest body dto.UpdatePhotoRequest true "Update photo body"
// @Success 200 {object}  helpers.ApiResponse{data=dto.UpdatePhotoResponse} "Success"
// @Failure 422 {object}  helpers.ApiResponse
// @Param photoId path uint true "Photo Id"
// @Router /photos/{photoId} [put]
func (controller *PhotoHandler) UpdatePhoto(ctx *gin.Context) {

	ids := ctx.Param("photoId")
	idconvert, errconvert := strconv.Atoi(ids)
	if errconvert != nil {
		response := helpers.GenerateApiResponse(http.StatusBadGateway, "Unable to parse id", nil)
		ctx.JSON(http.StatusBadGateway, response)
	}

	var updatePhotoRequest dto.UpdatePhotoRequest

	ctx.ShouldBindJSON(&updatePhotoRequest)

	errValid := helpers.CheckValid(updatePhotoRequest)
	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
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

// @Summary      Delete photo
// @Description  Delete photo
// @Tags         Photo
// @Security Authorization
// @Produce json
// @Success 200 {object}  helpers.ApiResponse{data=dto.DeletePhotoResponse} "Success"
// @Failure 422 {object}  helpers.ApiResponse
// @Param photoId path uint true "Photo Id"
// @Router /photos/{photoId} [delete]
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
