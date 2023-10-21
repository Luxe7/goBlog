package v1

import (
	"github.com/gin-gonic/gin"
	"goBlog/model"
	"goBlog/utlis/errormsg"
	"net/http"
	"strconv"
)

var code int

// Find User Exist
func UserExist(c *gin.Context) {

}

// Add User
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.UserName)
	if code == errormsg.SUCCESS {
		model.CreateUser(&data)
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

// Find User List
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
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

// Edit User
func EditUser(c *gin.Context) {

}

// Delete User
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})

}
