package router

import (
	"BookStore/model"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func BookAddBook(c *gin.Context) {
	title := c.PostForm("title")
	price, _ := com.StrTo(c.PostForm("price")).Float64()
	author := c.PostForm("author")
	sales, _ := com.StrTo(c.PostForm("sales")).Int()
	stock, _ := com.StrTo(c.PostForm("stock")).Int()
	imgPath := "/static/img/default.jpg"
	_ = model.AddBook(title, author, price, sales, stock, imgPath)
	c.Redirect(301, "/book/manager")
}
