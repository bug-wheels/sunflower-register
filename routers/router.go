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
		// 获取当前的服务信息
		serviceDiscoveryGroup.GET("/info", GetAllServiceInstance)
		serviceDiscoveryGroup.GET("/instances/:dc/:ns/:serviceId", GetServiceInstance)
		// 注册一个服务实例
		serviceDiscoveryGroup.POST("/register/:dc/:ns", RegisterServiceInstance)
		// 剔除一个服务实例
		serviceDiscoveryGroup.DELETE("/deregister/:dc/:ns/:instanceId", DeregisterServiceInstance)
	}
	return r
}
