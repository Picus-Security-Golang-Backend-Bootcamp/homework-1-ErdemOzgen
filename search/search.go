package search

import (
	"fmt"
	"homework-1-ErdemOzgen/model"
	"homework-1-ErdemOzgen/utils"
)

type Book model.Book

func SearchByAuthor(author string, books []Book) {
	counter := 0
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for _, book := range books {
		if book.Author == author {
			counter++
			fmt.Println("FOUNDED Count :", counter) // TODO: After location set get rid of this line
			//fmt.Println(book)
			utils.PrintPretty(&book)
		}

	}
	if counter == 0 {
		fmt.Println("NOT FOUND") // TODO: After location set
		counter = 0              // Not needed for now but will be useful later
	}

}
func SearchByTitle(title string, books []Book) {
	counter := 0
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for _, book := range books {
		if book.Title == title {
			counter++
			fmt.Println("FOUNDED Count :", counter) // TODO: After location set get rid of this line
			//fmt.Println(book)
			utils.PrintPretty(&book)
		}
	}
	if counter == 0 {
		fmt.Println("NOT FOUND") // TODO: After location set
		counter = 0              // Not needed for now but will be useful later
	}
}
func ListBooks(books []Book) {
	for _, book := range books {
		utils.PrintPretty(&book)
	}
}
