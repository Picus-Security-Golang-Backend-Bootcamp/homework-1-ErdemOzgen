package main

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/jsonops"
	"homework-1-ErdemOzgen/model"
	"homework-1-ErdemOzgen/utils"
)

type Book model.Book

func searchByAuthor(author string, books []Book) {
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for _, book := range books {
		if book.Author == author {
			fmt.Println("Founded") // TODO: After location set get rid of this line
			//fmt.Println(book)
			utils.PrintPretty(&book)
		}
	}

}

func searchByTitle(title string, books []Book) {
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for _, book := range books {
		if book.Title == title {
			fmt.Println("Founded") // TODO: After location set get rid of this line
			//fmt.Println(book)
			utils.PrintPretty(&book)
		}
	}
}

func listBooks(books []Book) {
	for _, book := range books {
		utils.PrintPretty(&book)
	}
}
func main() {
	fmt.Println("Staring Program")

	// ----------------------------------------------------------------
	// 1. Read json file and convert it to string
	// 2. Add them to slice
	var books []Book // slice of Books structs

	json.Unmarshal([]byte(jsonops.OpenJsonFile("test.json")), &books)
	fmt.Println(len(books))
	fmt.Println(books[0].Author)

	//----------------------------------------------------------------

	searchByAuthor("Erdem Ozgen", books)
	searchByTitle("Golang Book", books)
	searchByAuthor("Leo Tolstoy", books)
	//----------------------------------------------------------------
	// 3. Print all books
	//listBooks(books)

}
