package search

import (
	"fmt"
	"homework-1-ErdemOzgen/model"
	"homework-1-ErdemOzgen/utils"
)

type Book model.Book

func SearchByAuthor(author string, books []Book) {
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for i, book := range books {
		if book.Author == author {
			fmt.Println("Founded") // TODO: After location set get rid of this line
			//fmt.Println(book)
			utils.PrintPretty(&book)
		}
		if i == len(books)-1 && book.Author != author {
			fmt.Println("Not found") // TODO: After location set get rid of this line
		}
	}

}
func SearchByTitle(title string, books []Book) {
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for i, book := range books {
		if book.Title == title {
			fmt.Println("Founded") // TODO: After location set get rid of this line
			//fmt.Println(book)
			utils.PrintPretty(&book)
		}
		if i == len(books)-1 && book.Title != title {
			fmt.Println("Not found") // TODO: After location set get rid of this line
		}
	}
}
func ListBooks(books []Book) {
	for _, book := range books {
		utils.PrintPretty(&book)
	}
}
