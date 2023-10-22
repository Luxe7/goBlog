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
		routerV1.POST("/category/add", v1.AddCategory)
		routerV1.GET("/category", v1.GetCategory)
		routerV1.PUT("/category/:id", v1.EditCategory)
		routerV1.DELETE("/category/:id", v1.DeleteCategory)
		// Article
		routerV1.POST("/article/add", v1.AddArticle)
		routerV1.GET("/article/:id", v1.GetArticle)
		routerV1.GET("/article", v1.GetArticleList)
		routerV1.GET("/article/list", v1.GetArticlesInCategory)
		routerV1.PUT("/article/:id", v1.EditArticle)
		routerV1.DELETE("/article/:id", v1.DeleteArticle)
	}
	r.Run(utlis.HttpPort)

}
