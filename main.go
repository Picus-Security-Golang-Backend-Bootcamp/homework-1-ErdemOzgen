package main

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/jsonops"

	"github.com/kr/pretty"
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
	fmt.Println("Searching")
	for _, book := range books {
		if book.Author == author {
			fmt.Println("Founded")
			fmt.Println(book)
		}
	}

}

func searchByTitle(title string, books []Book) {
	fmt.Println("Searching")
	for _, book := range books {
		if book.Title == title {
			fmt.Println("Founded")
			fmt.Println(book)
			fmt.Printf("%# v", pretty.Formatter(book))
		}
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

}
