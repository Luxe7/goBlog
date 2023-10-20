package routes

import (
	"github.com/gin-gonic/gin"
	"goBlog/utlis"
)

func InitRouter() {
	gin.SetMode(utlis.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "Hello,World",
			})
		})

	}
	r.Run(utlis.HttpPort)

}
