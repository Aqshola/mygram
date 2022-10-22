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

type CommentHandler struct {
	service service.CommentService
}

func NewCommentController(service service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (controller *CommentHandler) Route(route *gin.Engine) {
	comment := route.Group("/comments").Use(middlewares.Authentication())
	comment.POST("/", controller.CreateComment)
	comment.GET("/", controller.GetAllComment)
	comment.Use(middlewares.Authorization("comments", "commentId")).PUT("/:commentId", controller.UpdateComment)
	comment.Use(middlewares.Authorization("comments", "commentId")).DELETE("/:commentId", controller.DeleteComment)

}

func (controller *CommentHandler) CreateComment(ctx *gin.Context) {
	var createCommentRequest model.CreateCommentRequest
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	ctx.ShouldBindJSON(&createCommentRequest)

	errValid := helpers.CheckValid(createCommentRequest)
	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res, errCreate := controller.service.CreateComment(userId, &createCommentRequest)
	if errCreate != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errCreate.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success add comment", res)
	ctx.JSON(http.StatusOK, response)
}

func (controller *CommentHandler) GetAllComment(ctx *gin.Context) {

	res, errGet := controller.service.GetAllComment()
	if errGet != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errGet.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success get all comment", res)
	ctx.JSON(http.StatusOK, response)
}

func (controller *CommentHandler) UpdateComment(ctx *gin.Context) {
	var updateCommentRequest model.UpdateCommentRequest
	ids := ctx.Param("commentId")
	idconvert, errconvert := strconv.Atoi(ids)
	if errconvert != nil {
		response := helpers.GenerateApiResponse(http.StatusBadGateway, "Unable to parse id", nil)
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	ctx.ShouldBindJSON(&updateCommentRequest)
	errValid := helpers.CheckValid(updateCommentRequest)
	if errValid != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errValid.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res, errDetail := controller.service.UpdateComment(uint(idconvert), &updateCommentRequest)
	if errDetail != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, errDetail.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success update comment", res)
	ctx.JSON(http.StatusOK, response)
}

func (controller *CommentHandler) DeleteComment(ctx *gin.Context) {

	ids := ctx.Param("commentId")
	idconvert, errconvert := strconv.Atoi(ids)
	if errconvert != nil {
		response := helpers.GenerateApiResponse(http.StatusBadGateway, "Unable to parse id", nil)
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	res, err := controller.service.DeleteComment(uint(idconvert))
	if err != nil {
		response := helpers.GenerateApiResponse(http.StatusUnprocessableEntity, err.Error(), nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.GenerateApiResponse(http.StatusOK, "Success delete photo", res)
	ctx.JSON(http.StatusOK, response)
}
