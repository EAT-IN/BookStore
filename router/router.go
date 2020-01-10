package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	// 加载静态文件
	r.Static("/static", "./static")
	// 制定index页面, 当一个目录下面有多个html文件时,不一样的层级需要单独拉出来
	r.LoadHTMLGlob("./template/*")

	r.GET("/", Index)
	r.GET("/login", Login)
	r.POST("/user/login", UserLogin)
	r.GET("/regist", Regist)
	r.POST("/user/regist", UserRegist)
	r.GET("/test", Test)
	r.GET("/manager", Manager)
	r.GET("/book/manager", BookManager)
	r.GET("/book/edit", BookEdit)
	r.GET("/book/addbook", BookAddBook)

	return r

}
