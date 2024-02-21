package router

import (
	"egardev-gin-socmed/config"
	"egardev-gin-socmed/handler"
	"egardev-gin-socmed/repository"
	"egardev-gin-socmed/service"
	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/register", authHandler.Register)
}