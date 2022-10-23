package handlers

import (
	"mygram/dto"
	"mygram/helpers"
	"mygram/middlewares"
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

// @Summary      Register User
// @Description Register new user
// @Tags         User
// @Accept json
// @Produce json
// @Param registerRequest body dto.RegisterRequest true "Register body"
// @Success 201 {object}  helpers.ApiResponse{data=dto.RegisterResponse} "Success"
// @Failure 422 {object}  helpers.ApiResponse
// @Router /users/register [post]
func (controller *UserHandler) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest

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

	response := helpers.GenerateApiResponse(http.StatusOK, "Success Register", res)
	ctx.JSON(http.StatusCreated, response)
}

// @Summary      Login
// @Description Login user
// @Tags         User
// @Accept json
// @Produce json
// @Param registerRequest body dto.LoginRequest true "Login body"
// @Success 200 {object}  helpers.ApiResponse{data=dto.LoginResponse} "Success"
// @Failure 422 {object}  helpers.ApiResponse
// @Router /users/login [post]
func (controller *UserHandler) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest

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

// @Summary     Update User
// @Description Update User Data
// @Tags         User
// @Security Authorization
// @Accept json
// @Produce json
// @Param registerRequest body dto.UpdateRequest true "Login body"
// @Success 200 {object}  helpers.ApiResponse{data=dto.UpdateResponse} "Success"
// @Failure 422 {object}  helpers.ApiResponse
// @Router /users/ [put]
func (controller *UserHandler) UpdateUser(ctx *gin.Context) {
	var updateRequest dto.UpdateRequest
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	ctx.ShouldBindJSON(&updateRequest)

	errValid := helpers.CheckValid(updateRequest)
	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
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

// @Summary     Delete User
// @Description Delete User
// @Tags         User
// @Security Authorization
// @Accept json
// @Produce json
// @Success 200 {object}  helpers.ApiResponse{data=dto.DeleteUserResponse} "Success"
// @Failure 422 {object}  helpers.ApiResponse
// @Router /users/ [delete]
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
