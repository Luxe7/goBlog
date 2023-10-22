package v1

import (
	"github.com/gin-gonic/gin"
	"goBlog/model"
	"goBlog/utlis/errormsg"
	"net/http"
	"strconv"
)

// Add Category
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errormsg.SUCCESS {
		model.CreateCategory(&data)
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

// Find Category List
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCategory(pageSize, pageNum)
	if data != nil {
		code = errormsg.SUCCESS
	} else {
		code = errormsg.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})

}

// Find Category Articles
// Edit Category
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id")) //为什么这个地方的id是从Param中获得的，不可以使用Data.ID呢
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errormsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	if code == errormsg.SUCCESS {
		model.UpdateCategory(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

// Delete Category
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})

}
