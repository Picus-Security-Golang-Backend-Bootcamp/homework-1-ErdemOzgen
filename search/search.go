package search

import (
	"fmt"
	"homework-1-ErdemOzgen/model"
	"homework-1-ErdemOzgen/utils"
)

type Book model.Book

func SearchByAuthor(author string, books []Book) {
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for _, book := range books {
		if book.Author == author {
			fmt.Println("Founded") // TODO: After location set get rid of this line
			//fmt.Println(book)
			utils.PrintPretty(&book)
		}
	}

}
func SearchByTitle(title string, books []Book) {
	fmt.Println("Searching...") // TODO: After location set get rid of this line
	for _, book := range books {
		if book.Title == title {
			fmt.Println("Founded") // TODO: After location set get rid of this line
			//fmt.Println(book)
			utils.PrintPretty(&book)
		}
	}
}
