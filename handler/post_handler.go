package handler

import (
	"egardev-gin-socmed/dto"
	"egardev-gin-socmed/errorHandler"
	"egardev-gin-socmed/helper"
	"egardev-gin-socmed/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *postHandler {
	return &postHandler{
		service: s,
	}
}

func (h postHandler) Create(c *gin.Context) {
	var post dto.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		errorHandler.HandleError(c, &errorHandler.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("public/picture", 0755); err != nil {
			errorHandler.HandleError(c, &errorHandler.InternalServerError{Message: err.Error()})
			return
		}

		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)
	}

	userId := 1
	post.UserID = userId

	if err := h.service.Create(&post); err != nil {
		errorHandler.HandleError(c, err)
		return
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Success post your tweet",
	})

	c.JSON(http.StatusCreated, response)
}
