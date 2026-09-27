package router

import (
	"projeto-teste/handler"
	"github.com/gin-gonic/gin"
)
func InitProductRouter(router *gin.Engine, handler *handler.ProductHandler) {
	routes := router.Group("/api/products")

	routes.POST("", handler.Create)
	routes.GET("/:id", handler.Read)
	// routes.DELETE("/:id", handler.Delete)
}

