package routers

import "github.com/gin-gonic/gin"

type ZoneInfo struct {
	Datacenter string `uri:"dc" binding:"required" json:"datacenter"`
	Namespace  string `uri:"ns" binding:"required" json:"namespace"`
}

func RegisterServiceInstance(c *gin.Context) {
	var zoneInfo ZoneInfo

	if err := c.ShouldBindUri(&zoneInfo); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, zoneInfo)
}
