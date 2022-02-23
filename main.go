package main

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/jsonops"
	"homework-1-ErdemOzgen/restapi"
	"homework-1-ErdemOzgen/search"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	//----------------------------------------------------------------
	// 4. Add extra REST API

	router := mux.NewRouter()
	router.HandleFunc("/books", restapi.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{author}", restapi.GetBookByAuthor).Methods("GET")
	router.HandleFunc("/books/{title}", restapi.GetBookByTitle).Methods("GET")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))

}
