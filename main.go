package main

import (
	"BookStore/model"
	"BookStore/router"
	"fmt"
)

func main() {
	// 初始化数据库
	model.InitDB()
	// 初始化路由设置
	r := router.InitRouter()
	// 启动服务器
	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
	}

}
