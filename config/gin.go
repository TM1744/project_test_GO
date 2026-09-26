package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initGin() *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.Default())
    return engine
}