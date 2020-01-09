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
