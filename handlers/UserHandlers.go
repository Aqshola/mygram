package handlers

import (
	"mygram/helpers"
	"mygram/model"
	"mygram/service"
	"net/http"

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
	user.POST("/login")
	// user.Use(middlewares.Authentication()).PUT("/:userId", UpdateUser)
	// user.Use(middlewares.Authentication()).DELETE("/:userId",DeleteUser)

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
	return

}

func (controller *UserHandler) Login(ctx *gin.Context) {

}

func (controller *UserHandler) UpdateUser(ctx *gin.Context) {

}

func (controller *UserHandler) DeleteUser(ctx *gin.Context) {

}
