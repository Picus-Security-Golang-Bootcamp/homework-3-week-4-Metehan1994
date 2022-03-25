package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name              string
	NumOfPages        int
	NumOfBooksInStock int
	Price             int
	StockCode         string
	ISBN              string
	AuthorName        string
	Author            Author `gorm:"foreignKey:Name;references:AuthorName"`
}

func (Book) TableName() string {
	return "Book"
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, Pages: %d, Price: %d, ISBN: %s, AuthorName: %s, CreatedAt : %s",
		b.ID, b.Name, b.NumOfPages, b.Price, b.ISBN, b.AuthorName, b.CreatedAt.Format("2006-01-02 15:04:05"))
}
