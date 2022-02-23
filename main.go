package main

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/jsonops"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Pages     int    `json:"pages"`
	Country   string `json:"country"`
	ImageLink string `json:"imageLink"`
	Year      string `json:"year"`
	Language  string `json:"language"`
	Link      string `json:"link"`
}

func searchByAuthor(author string, books []Book) {
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for _, book := range books {
		if book.Author == author {
			fmt.Println("Founded") // TODO: After location set get rid of this line
			//fmt.Println(book)
			printPretty(&book)
		}
	}

}
func printPretty(book interface{}) {
	b, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}

func searchByTitle(title string, books []Book) {
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for _, book := range books {
		if book.Title == title {
			fmt.Println("Founded") // TODO: After location set get rid of this line
			//fmt.Println(book)
			printPretty(&book)
		}
	}
}

func listBooks(books []Book) {
	for _, book := range books {
		printPretty(&book)
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
	//----------------------------------------------------------------
	// 3. Print all books
	listBooks(books)

}
