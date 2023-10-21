package routes

import (
	"github.com/gin-gonic/gin"
	v1 "goBlog/api/v1"
	"goBlog/utlis"
)

func InitRouter() {
	gin.SetMode(utlis.AppMode)
	r := gin.Default()
	routerV1 := r.Group("api/v1")
	{
		// USer
		routerV1.POST("/user/add", v1.AddUser)
		routerV1.GET("/user/get", v1.GetUsers)
		routerV1.PUT("/user/:id", v1.EditUser)
		routerV1.DELETE("/user/:id", v1.DeleteUser)

		// Category
		// Article
	}
	r.Run(utlis.HttpPort)

}
