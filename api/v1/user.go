package v1

import (
	"go-project-back/model"
	"go-project-back/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

// 添加用户
func AddUser(c *gin.Context){
	// todo 添加用户
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}
// 查询用户列表
func GetUsers(c *gin.Context){
	// 这里的 pagesize, pagenum 就是请求里面带着的
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}

	data := model.GetUsers(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"statues": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}
// 编辑用户
func EditUser(c *gin.Context){
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		// 终止
		c.Abort()
	}

	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
// 删除用户
func DeleteUser(c *gin.Context){
	// 注意这里数据库都是软删除，数据库里还是有显示，但是查询不出来
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}