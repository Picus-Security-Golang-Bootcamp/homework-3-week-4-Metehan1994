package book

import (
	"github.com/Metehan1994/HWs/HW3/domain/author"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookName          string
	NumOfPages        int
	NumOfBooksinStock int
	Price             int
	StockCode         string
	ISBN              string
	AuthorName        string
	Author            []author.Author `gorm:"foreignKey:AuthorName;references:AuthorName"`
}

func (Book) TableName() string {
	return "Book"
}
