package main

import (
	"go-project-back/model"
	"go-project-back/routes"
)

func main() {

	// 引用数据库

	model.InitDb()
	routes.InitRouter()
}