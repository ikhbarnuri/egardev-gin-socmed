package router

import (
	"egardev-gin-socmed/config"
	"egardev-gin-socmed/handler"
	"egardev-gin-socmed/repository"
	"egardev-gin-socmed/service"
	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/tweets")

	r.POST("/", postHandler.Create)
}
