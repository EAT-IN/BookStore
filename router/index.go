package router

import (
	"BookStore/model"
	"github.com/gin-gonic/gin"
	"html/template"
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

func Manager(c *gin.Context) {
	//后台管理页面
	c.HTML(200, "manager.html", nil)

}

func BookManager(c *gin.Context) {
	// 图书管理页面
	books := model.GetBooks()
	c.HTML(200, "book_manager.html", books)

}

func BookEdit(c *gin.Context) {
	c.HTML(200, "book_edit.html", nil)

}

// 渲染模版样例
func Test(c *gin.Context) {
	data := `<table border="1">
			<tr>
				<td>row 1, cell 1</td>
				<td>row 1, cell 2</td>
			</tr>
			<tr>
				<td>row 2, cell 1</td>
				<td>row 2, cell 2</td>
			</tr>
			</table>`
	c.HTML(200, "demo.html", template.HTML(data))
}
