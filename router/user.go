package router

import (
	"BookStore/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := model.CheckUserNameAndPassword(username, password)
	if err != nil {
		// 密码错误或者用户不存在， 则返回login页面 继续登录
		//fmt.Println(err.Error())
		c.HTML(200, "login.html", err.Error())
	} else {
		fmt.Println(user)
		c.HTML(200, "login_success.html", nil)
	}
}

func UserRegist(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	err := model.SaveUser(username, password, email)
	if err != nil {
		c.HTML(200, "regist.html", err.Error())
	} else {
		c.HTML(200, "regist_success.html", nil)
	}

}
