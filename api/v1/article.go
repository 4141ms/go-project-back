package v1

import (
	// "fmt"
	"go-project-back/model"
	"go-project-back/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(c *gin.Context){
	// todo 添加分类
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo 查询分类下的所有文章

// todo 查询单个文章信息


// todo 查询文章列表
func GetArt(c *gin.Context){
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
// 编辑文章名
func EditArt(c *gin.Context){
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	model.EditArt(id, &data)

	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
// 删除文章
func DeleteArt(c *gin.Context){
	// 注意这里数据库都是软删除，数据库里还是有显示，但是查询不出来
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}