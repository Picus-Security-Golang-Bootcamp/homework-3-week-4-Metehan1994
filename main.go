package main

import (
	"fmt"
	"log"

	postgres "github.com/Metehan1994/HWs/HW3/common/db"
	"github.com/Metehan1994/HWs/HW3/csv_utils"
	"github.com/Metehan1994/HWs/HW3/domain/entities"
	"github.com/Metehan1994/HWs/HW3/domain/repos"

	"github.com/joho/godotenv"
)

var filename string = "book.csv"

func main() {
	//CSV to book struct
	bookList, err := csv_utils.ReadCSV(filename)
	if err != nil {
		log.Fatal(err)
	}

	//Set environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Creating DB connection with postgres
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	//Author Repo
	authorRepo := repos.NewAuthorRepository(db)
	authorRepo.Migrations()
	authorRepo.InsertSampleData(bookList)

	//Book Repo
	bookRepo := repos.NewBookRepository(db)
	bookRepo.Migrations()
	bookRepo.InsertSampleData(bookList)

	QueriesforRepos(authorRepo, bookRepo)
}

func QueriesforRepos(authorRepo *repos.AuthorRepository, bookRepo *repos.BookRepository) {

	//Queries for Book repo

	bookRepo.List() //ListBooks

	book, err := bookRepo.GetByID(2) //GetBookByID
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(book.ToString())

	bookRepo.FindByWord("AND") //FindBookByWord case insensitive

	bookRepo.FindByName("War and Peace") //FindBookByWholeName case sensitive

	bookRepo.DeleteByName("Anna Karenina") //DeleteBookByName

	bookRepo.DeleteById(5) //DeleteBookByID

	books, err := bookRepo.GetBooksWithAuthorInformation() ////Get book with author
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		fmt.Println(book)
	}

	err = bookRepo.Buy(12, 7) //BuyBookByID
	if err != nil {
		fmt.Println(err)
	}

	bookRepo.MaxPrice() //Finding most expensive book

	bookRepo.PriceBetweenFromLowerToUpper(15, 25) //Finding books inside a price range

	/***********/

	//Queries for Author repo

	authorRepo.List() //ListAuthors

	author, err := authorRepo.GetByID(3) //GetAuthorByID
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(author.ToString())

	authorRepo.FindByWord("EL") //FindAuthorByWord- case insensitive

	authorRepo.FindByName("Lev Tolstoy") //FindBookByWholeName - case sensitive

	authorRepo.DeleteByName("Lev Tolstoy") //DeleteAuthorByName

	authorRepo.DeleteById(8) //DeleteAuthorByID

	authors, err := authorRepo.GetAuthorsWithBookInformation() //Get author with book
	if err != nil {
		log.Fatal(err)
	}
	for _, author := range authors {
		fmt.Println(author)
	}

	authorRepo.BooksOfAuthors("Fyodor Dostoyevski") //Getting books of an author

	//Create new author and book together
	newAuthor := entities.Author{Name: "Franz Kafka", ID: 9}
	authorRepo.Create(newAuthor)
	newBook := entities.Book{Name: "The Metamorphosis", NumOfPages: 100,
		NumOfBooksInStock: 10, Price: 30, StockCode: "Book1121", ISBN: "9780553213690", AuthorID: 9}
	bookRepo.Create(newBook)

}
