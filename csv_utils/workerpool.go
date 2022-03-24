package csv_utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/Metehan1994/HWs/HW3/models"
)

func ReadCSVWithWorkerPool(filename string) error {
	numJobs := 10
	jobs := make(chan []string, numJobs)
	results := make(chan models.Book, numJobs)

	wg := sync.WaitGroup{}

	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go toStruct(jobs, results, &wg)
	}

	go func() {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		csvReader := csv.NewReader(file)
		lines, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		for _, line := range lines[1:] {
			jobs <- line
		}

		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for v := range results {
		fmt.Println(v)
	}

	return nil
}

func toStruct(jobs <-chan []string, results chan<- models.Book, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		book := models.Book{}
		book.BookName = j[0]
		book.NumOfPages, _ = strconv.Atoi(j[1])
		book.NumOfBooksinStock, _ = strconv.Atoi(j[2])
		book.Price, _ = strconv.Atoi(j[3])
		book.StockCode = j[4]
		book.ISBN = j[5]
		book.Author.AuthorName = j[6]

		results <- book
	}
}