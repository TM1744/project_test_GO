package router

import (
	"projeto-teste/handler"

	"github.com/gin-gonic/gin"
)

func InitAuditRouter(router *gin.Engine, handler *handler.AuditHandler) {
	routes := router.Group("/api/audits")

	routes.GET("", handler.ReadAll)
	routes.GET("/:id", handler.Read)
}