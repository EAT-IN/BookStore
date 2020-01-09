package router

import (
	"BookStore/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func BookManager(c *gin.Context) {
	// 图书管理页面
	books := model.GetBooks()
	fmt.Println(books)
	c.HTML(200, "book_manager.html", books)

}
