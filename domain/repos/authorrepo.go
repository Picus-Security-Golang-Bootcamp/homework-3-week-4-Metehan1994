package repos

import (
	"errors"
	"fmt"

	"github.com/Metehan1994/HWs/HW3/domain/entities"
	"github.com/Metehan1994/HWs/HW3/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (a *AuthorRepository) List() {
	var authors []entities.Author
	a.db.Find(&authors)

	for _, author := range authors {
		fmt.Println(author.ToString())
	}
}

func (a *AuthorRepository) GetByID(ID int) (*entities.Author, error) {
	var author entities.Author
	result := a.db.First(&author, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &author, nil
}

func (a *AuthorRepository) FindByName(name string) {
	var authors []entities.Author
	a.db.Where("name LIKE ? ", "%"+name+"%").Find(&authors)

	for _, author := range authors {
		fmt.Println(author.ToString())
	}
}

func (a *AuthorRepository) Create(name string) error {
	var author entities.Author
	author.Name = name
	result := a.db.Create(&author)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *AuthorRepository) DeleteByName(name string) error {
	result := a.db.Where("name = ?", name).Delete(&entities.Author{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *AuthorRepository) DeleteById(id int) error {
	result := a.db.Delete(&entities.Author{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func (a *AuthorRepository) GetAuthorsWithBookInformation() ([]entities.Author, error) {
// 	var authors []entities.Author
// 	result := a.db.Preload("Book").Find(&authors)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return authors, nil
// }

func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&entities.Author{})
}

func (a *AuthorRepository) InsertSampleData(bookList models.BookList) {

	authors := []entities.Author{}
	for _, book := range bookList {
		newAuthor := entities.Author{
			Name: book.Author.AuthorName,
		}
		authors = append(authors, newAuthor)
	}

	for _, author := range authors {
		a.db.Where(entities.Author{Name: author.Name}).FirstOrCreate(&author)
	}

}
