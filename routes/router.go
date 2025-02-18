package routes

import (
	v1 "go-project-back/api/v1"
	"go-project-back/utils"

	"github.com/gin-gonic/gin"
)

// 方法首字母大写表示 这是一个全局方法
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 测试接口 localhost:3000/api/v1/hello
		// routerV1.GET("hello", func(c *gin.Context){
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"msg": "ok",
		// 	})
		// })


		// 用户模块的路由接口
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUsers)
		router.PUT("user/:id",v1.EditUser)
		router.DELETE("user/:id",v1.DeleteUser)
		// 分类模块的用户接口
		router.POST("category/add",v1.AddCategory)
		router.GET("category",v1.GetCate)
		router.PUT("category/:id",v1.EditCate)
		router.DELETE("category/:id",v1.DeleteCate)
		// 文章模块的用户接口
		router.POST("article/add",v1.AddArticle)
		router.GET("article",v1.GetArt)
		router.GET("article/list/:id",v1.GetCateArt)
		router.GET("article/info/:id",v1.GetArtInfo)
		router.PUT("article/:id",v1.EditArt)
		router.DELETE("article/:id",v1.DeleteArt)
	}

	r.Run(utils.HttpPort)
}