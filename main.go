package main

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/jsonops"
	"homework-1-ErdemOzgen/search"
)

//type Book model.Book // Questtion []search.Book is good implementation???

func main() {

	fmt.Println("ASCII ART")
	fmt.Println("Staring Program...")

	var books []search.Book // slice of Books structs

	// ----------------------------------------------------------------
	// 1. Read json file and convert it to string
	// 2. Add them to slice

	json.Unmarshal([]byte(jsonops.OpenJsonFile("test.json")), &books)
	fmt.Println(len(books))
	fmt.Println(books[0].Author)

	//----------------------------------------------------------------

	search.SearchByAuthor("Erdem Ozgen", books)
	search.SearchByTitle("Golang Book", books)
	//search.SearchByAuthor("Leo Tolstoy", books)
	//----------------------------------------------------------------
	// 3. Print all books
	//listBooks(books)

}
