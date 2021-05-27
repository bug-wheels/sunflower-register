package routers

import (
	"github.com/gin-gonic/gin"
	"sunflower/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	serviceDiscoveryGroup := r.Group("/api/v1/service")
	{
		serviceDiscoveryGroup.POST("/register/:dc/:ns", RegisterServiceInstance)
	}
	return r
}
