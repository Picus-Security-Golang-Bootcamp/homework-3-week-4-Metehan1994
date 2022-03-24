package book

import (
	"github.com/Metehan1994/HWs/HW3/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
}

func (b *BookRepository) InsertSampleData(bookList models.BookList) {

	books := []Book{}
	for _, book := range bookList {
		newBook := Book{
			BookName:          book.BookName,
			NumOfPages:        book.NumOfPages,
			NumOfBooksinStock: book.NumOfBooksinStock,
			Price:             book.Price,
			StockCode:         book.StockCode,
			ISBN:              book.ISBN,
			AuthorName:        book.Author.AuthorName,
		}
		books = append(books, newBook)
	}

	for _, eachBook := range books {
		b.db.Create(&eachBook)
	}

}
