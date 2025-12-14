package routes

import (
	"astrodev.online/multimedia-server/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUploadRoutes(router *gin.Engine, handler *handlers.UploadHandler){
	route := router.Group("/api/v1/image")
	route.POST("/single",handler.UploadSingleImage)
}


