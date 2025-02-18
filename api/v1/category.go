package v1

import (
	"go-project-back/model"
	"go-project-back/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var code int

// 添加分类
func AddCategory(c *gin.Context){
	// todo 添加分类
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		code = errmsg.ERROR_CATEGORY_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo 查询分类下的所有文章


// 查询分类列表
func GetCate(c *gin.Context){
	// 这里的 pagesize, pagenum 就是请求里面带着的
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}

	data := model.GetCategory(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"statues": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}
// 编辑分类名
func EditCate(c *gin.Context){
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		// 终止
		c.Abort()
	}

	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
// 删除分类
func DeleteCate(c *gin.Context){
	// 注意这里数据库都是软删除，数据库里还是有显示，但是查询不出来
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}