package router

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// 主页面
	c.HTML(200, "index.html", nil)

}

func Login(c *gin.Context) {
	// 登录页面
	c.HTML(200, "login.html", nil)

}

func Regist(c *gin.Context) {
	// 注册页面
	c.HTML(200, "regist.html", nil)

}
