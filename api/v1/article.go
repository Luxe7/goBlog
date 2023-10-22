package v1

import (
	"github.com/gin-gonic/gin"
	"goBlog/model"
	"goBlog/utlis/errormsg"
	"net/http"
	"strconv"
)

// Add Article
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CheckArticle(data.Title)
	if code == errormsg.SUCCESS {
		model.CreateArticle(&data)
	}
	if code == errormsg.ERROR_USERNAME_USED {
		code = errormsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

// Find Articles in Category
func GetArticlesInCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArticlesInCategory(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

// Find one Article
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

// Find ArticleList
func GetArticleList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArticleList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

// Find Article List
// Edit Article
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id")) //为什么这个地方的id是从Param中获得的，不可以使用Data.ID呢
	_ = c.ShouldBindJSON(&data)
	code := model.CheckArticle(data.Title)
	if code == errormsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	if code == errormsg.SUCCESS {
		model.UpdateArticle(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

// Delete Article
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})

}
