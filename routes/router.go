package routes

import (
	"go-project-back/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 方法首字母大写表示 这是一个全局方法
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}