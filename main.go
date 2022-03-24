package main

import (
	"log"

	postgres "github.com/Metehan1994/HWs/HW3/common/db"
	csv_utils "github.com/Metehan1994/HWs/HW3/csv_utils"
	"github.com/joho/godotenv"
)

var filename string = "book.csv"

func main() {
	//CSV to book struct
	err := csv_utils.ReadCSVWithWorkerPool(filename)
	if err != nil {
		log.Fatal(err)
	}

	//Set environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Creating DB connection with postgres
	_, err = postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")
}
