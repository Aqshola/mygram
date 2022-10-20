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
	user.Use(middlewares.Authentication()).PUT("/:userId", controller.UpdateUser)
	user.Use(middlewares.Authentication()).DELETE("/:userId", controller.DeleteUser)

}

func (controller *UserHandler) Register(ctx *gin.Context) {
	var registerRequest model.RegisterRequest

	ctx.ShouldBindJSON(&registerRequest)

	valid, trans := helpers.Valid()
	err := valid.Struct(registerRequest)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, e.Translate(trans), nil)
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}
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
	valid, trans := helpers.Valid()
	err := valid.Struct(loginRequest)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, e.Translate(trans), nil)
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}
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
	ids := ctx.Param("userId")

	idconvert, errconvert := strconv.Atoi(ids)
	if errconvert != nil {
		response := helpers.GenerateApiResponse(http.StatusBadGateway, "Unable to parse id", nil)
		ctx.JSON(http.StatusBadGateway, response)
	}

	ctx.ShouldBindJSON(&updateRequest)
	valid, trans := helpers.Valid()
	err := valid.Struct(updateRequest)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, e.Translate(trans), nil)
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	res, errUpdate := controller.service.UpdateUser(uint(idconvert), &updateRequest)
	if errUpdate != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errUpdate.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, "Update success", res)
	ctx.JSON(http.StatusOK, response)
}

func (controller *UserHandler) DeleteUser(ctx *gin.Context) {
	ids := ctx.Param("userId")

	idconvert, errconvert := strconv.Atoi(ids)
	if errconvert != nil {
		response := helpers.GenerateApiResponse(http.StatusBadGateway, "Unable to parse id", nil)
		ctx.JSON(http.StatusBadGateway, response)
	}

	errDelete := controller.service.DeleteUser(uint(idconvert))

	if errDelete != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errDelete.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Delete success", model.DeleteResponse{
		Message: "Your account has been successfully deleted",
	})

	ctx.JSON(http.StatusOK, response)
}
