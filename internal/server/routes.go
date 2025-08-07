package server

import (
	"ultahost-ai-gateway/internal/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(api.AuthMiddleware())

	r.POST("/chat", api.HandleChat)
}
