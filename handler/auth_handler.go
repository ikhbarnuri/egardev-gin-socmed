package handler

import (
	"egardev-gin-socmed/dto"
	"egardev-gin-socmed/errorHandler"
	"egardev-gin-socmed/helper"
	"egardev-gin-socmed/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorHandler.HandleError(c, &errorHandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorHandler.HandleError(c, &errorHandler.BadRequestError{Message: err.Error()})
		return
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register successfully, please login",
	})

	c.JSON(http.StatusCreated, response)
}
