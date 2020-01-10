package model

import (
	"time"
)

type Book struct {
	ID         int
	Title      string    `gorm:"type:varchar(100);NOT NULL;unique_index:uqi"`
	Author     string    `gorm:"type:varchar(100);NOT NULL;unique_index:uqi"`
	Price      float64   `gorm:"type:decimal(11,2);NOT NULL"`
	Sales      int       `gorm:"type:int(11);NOT NULL"`
	Stock      int       `gorm:"type:int(11);NOT NULL"`
	ImgPath    string    `gorm:"type:varchar(100);NOT NULL"`
	CreateTime time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"type:timestamp;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func GetBooks() []*Book {
	var books []*Book
	DB.Find(&books)
	return books

}

func AddBook(title string, author string, price float64, sales int, stock int, imgPath string) error {

	book := Book{
		Title:   title,
		Author:  author,
		Price:   price,
		Sales:   sales,
		Stock:   stock,
		ImgPath: imgPath,
	}
	err := DB.Create(&book).Error
	if err != nil {
		return err
	} else {
		return nil
	}

}
