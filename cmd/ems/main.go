package main

import (
	"ems-adming-go/internal/config"
	"ems-adming-go/internal/handler"
)

func main() {

	//	初始化数据库
	if err := config.InitDB(); err != nil {
		println("数据库初始化失败")
	}
	defer config.CloseDB()
	//	初始化路由
	router := handler.SetupRoutes()

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
