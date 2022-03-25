package main

import (
	"log"

	postgres "github.com/Metehan1994/HWs/HW3/common/db"
	"github.com/Metehan1994/HWs/HW3/csv_utils"
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

	//Queries for Book repo

	//bookRepo.List() //ListBooks

	// book, err := bookRepo.GetByID(2) //GetBookByID
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(book.ToString())

	//bookRepo.FindByWord("and") //FindBookByWord

	//bookRepo.FindByName("War and Peace") //FindBookByWholeName

	//bookRepo.DeleteByName("The Metamorphosis") //DeleteBookByName

	//bookRepo.DeleteById(5) //DeleteBookByID

	// books, err := bookRepo.GetBooksWithAuthorInformation() ////Get book with author
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, book := range books {
	// 	fmt.Println(book)
	// }

	// err = bookRepo.Buy(12, 7) //BuyBookByID
	// if err != nil {
	// 	fmt.Println(err)
	// }

	/***********/

	//Queries for Author repo

	//authorRepo.List() //ListAuthors

	// author, err := authorRepo.GetByID(3) //GetAuthorByID
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(author.ToString())

	//authorRepo.FindByWord("el") //FindAuthorByWord

	//authorRepo.FindByName("Lev Tolstoy") //FindBookByWholeName

	//authorRepo.DeleteByName("Lev Tolstoy") //DeleteAuthorByName

	//authorRepo.DeleteById(8) //DeleteAuthorByID

	// authors, err := authorRepo.GetAuthorsWithBookInformation() //Get author with book
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, author := range authors {
	// 	fmt.Println(author)
	// }
}
