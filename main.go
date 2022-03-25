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
	bookList, err := csv_utils.ReadCSVWithWorkerPool(filename)
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

	//Queries for Author repo

	//authorRepo.List() //ListAuthors

	// author, err := authorRepo.GetByID(10) //GetAuthorByID
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(author.ToString())

	// authorRepo.FindByName("el") //FindAuthorByName

	// authorRepo.Create("Seneca") //CreateNewAuthor

	//authorRepo.DeleteByName("Ahmet") //DeleteAuthorByName

	// authorRepo.DeleteById(8) //DeleteAuthorByID

	//Book Repo
	bookRepo := repos.NewBookRepository(db)
	bookRepo.Migrations()
	bookRepo.InsertSampleData(bookList)

	//Queries for Book repo

	//bookRepo.List() //ListBooks

	// book, err := bookRepo.GetByID(5) //GetBookByID
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(book.ToString())

	//bookRepo.FindByName("and") //FindBookByName

	//bookRepo.Create("The Metamorphosis", 104, 10, 30, "BOOK1121", "9789750719356", "Franz Kafka") //CreateNewBook

	// bookRepo.DeleteByName("The Metamorphosis") //DeleteBookByName

	//bookRepo.DeleteById(5) //DeleteBookByID

	// books, err := bookRepo.GetBooksWithAuthorInformation()
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
}
