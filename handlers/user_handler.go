package handlers

import (
	"mygram/helpers"
	"mygram/middlewares"
	"mygram/model"
	"mygram/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (controller *UserHandler) Route(route *gin.Engine) {
	user := route.Group("/users")
	user.POST("/register", controller.Register)
	user.POST("/login", controller.Login)
	user.Use(middlewares.Authentication()).PUT("/", controller.UpdateUser)
	user.Use(middlewares.Authentication()).DELETE("/", controller.DeleteUser)

}

func (controller *UserHandler) Register(ctx *gin.Context) {
	var registerRequest model.RegisterRequest

	ctx.ShouldBindJSON(&registerRequest)
	errValid := helpers.CheckValid(registerRequest)

	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	res, errRegister := controller.service.Register(&registerRequest)
	if errRegister != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errRegister.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, "Success Register", res)
	ctx.JSON(http.StatusCreated, response)
}

func (controller *UserHandler) Login(ctx *gin.Context) {
	var loginRequest model.LoginRequest

	ctx.ShouldBindJSON(&loginRequest)
	errValid := helpers.CheckValid(loginRequest)
	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
	}

	res, errLogin := controller.service.Login(&loginRequest)
	if errLogin != nil {
		response := helpers.GenerateApiResponse(http.StatusUnauthorized, errLogin.Error(), res)
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success Login", res)
	ctx.JSON(http.StatusOK, response)
}

func (controller *UserHandler) UpdateUser(ctx *gin.Context) {
	var updateRequest model.UpdateRequest
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	ctx.ShouldBindJSON(&updateRequest)

	errValid := helpers.CheckValid(updateRequest)
	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
	}

	res, errUpdate := controller.service.UpdateUser(userId, &updateRequest)
	if errUpdate != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errUpdate.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, "Update success", res)
	ctx.JSON(http.StatusOK, response)
}

func (controller *UserHandler) DeleteUser(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	resDelete, errDelete := controller.service.DeleteUser(uint(userId))

	if errDelete != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errDelete.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Delete success", resDelete)

	ctx.JSON(http.StatusOK, response)
}
