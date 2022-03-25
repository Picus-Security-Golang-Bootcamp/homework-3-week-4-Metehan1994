package repos

import (
	"errors"
	"fmt"

	"github.com/Metehan1994/HWs/HW3/domain/entities"
	"github.com/Metehan1994/HWs/HW3/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) List() {
	var books []entities.Book
	b.db.Find(&books)

	for _, author := range books {
		fmt.Println(author.ToString())
	}
}

func (b *BookRepository) GetByID(ID int) (*entities.Book, error) {
	var book entities.Book
	result := b.db.First(&book, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &book, nil
}

func (b *BookRepository) FindByName(name string) {
	var books []entities.Book
	b.db.Where("name LIKE ? ", "%"+name+"%").Find(&books)

	for _, author := range books {
		fmt.Println(author.ToString())
	}
}

func (b *BookRepository) DeleteByName(name string) error {
	result := b.db.Where("name = ?", name).Delete(&entities.Book{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) DeleteById(id int) error {
	result := b.db.Delete(&entities.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) Buy(quantity, id int) error {
	var book entities.Book
	result := b.db.First(&book, id)
	if result.Error != nil {
		return result.Error
	} else if book.NumOfBooksInStock < quantity {
		return fmt.Errorf("not Enough Book. Books: %d", book.NumOfBooksInStock)
	} else {
		fmt.Println("Successful Operation.")
	}

	result = b.db.Model(&book).Where("id = ? AND num_of_books_in_stock >= ?", id, quantity).
		Update("num_of_books_in_stock", gorm.Expr("num_of_books_in_stock - ?", quantity))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) GetBooksWithAuthorInformation() ([]entities.Book, error) {
	var books []entities.Book
	result := b.db.Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&entities.Book{})
}

func (b *BookRepository) InsertSampleData(bookList models.BookList) {

	books := []entities.Book{}
	for _, book := range bookList {
		newBook := entities.Book{
			Name:              book.BookName,
			NumOfPages:        book.NumOfPages,
			NumOfBooksInStock: book.NumOfBooksinStock,
			Price:             book.Price,
			StockCode:         book.StockCode,
			ISBN:              book.ISBN,
			AuthorID:          uint(book.Author.AuthorID),
		}
		books = append(books, newBook)
	}

	for _, eachBook := range books {
		b.db.Where(entities.Book{Name: eachBook.Name}).FirstOrCreate(&eachBook)
	}

}
