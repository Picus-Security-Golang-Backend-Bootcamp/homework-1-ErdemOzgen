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
	ImageLink string `json:"image_link"`
	Year      string `json:"year"`
	Language  string `json:"language"`
}

func main() {
	fmt.Println("Staring Program")

	// ----------------------------------------------------------------
	// 1. Read json file and convert it to string
	// 2. Add them to slice
	var books []Book // slice of Books structs

	json.Unmarshal([]byte(jsonops.OpenJsonFile("test.json")), &books)
	fmt.Println(len(books))
	// ----------------------------------------------------------------

}
