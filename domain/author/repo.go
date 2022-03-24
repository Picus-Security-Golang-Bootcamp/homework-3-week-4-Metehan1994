package author

import (
	"github.com/Metehan1994/HWs/HW3/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&Author{})
}

func (a *AuthorRepository) InsertSampleData(bookList models.BookList) {

	authors := []Author{}
	for _, book := range bookList {
		newAuthor := Author{
			AuthorName: book.Author.AuthorName,
		}
		authors = append(authors, newAuthor)
	}

	for _, author := range authors {
		a.db.Where(Author{AuthorName: author.AuthorName}).Attrs(Author{AuthorName: author.AuthorName}).FirstOrCreate(&author)
	}

}
