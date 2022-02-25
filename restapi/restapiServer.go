package main

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/jsonops"
	"homework-1-ErdemOzgen/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book model.Book

var books []Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, b := range books {
		//fmt.Println(params["author"])
		if b.Author == params["author"] {
			fmt.Println(b.Author)
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	json.NewEncoder(w).Encode("Book not found")
}

func GetBookT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, b := range books {
		//fmt.Println(params["author"])
		if b.Title == params["title"] {
			fmt.Println(b.Author)
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	json.NewEncoder(w).Encode("Book not found")
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var b Book

	_ = json.NewDecoder(r.Body).Decode(&b)

	books = append(books, b)
	json.NewEncoder(w).Encode(b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, b := range books {
		if b.Author == params["author"] {
			copy(books[i:], books[i+1:])
			books = books[:len(books)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	json.Unmarshal([]byte(jsonops.OpenJsonFile("test.json")), &books)

	router := mux.NewRouter()

	router.HandleFunc("/books", GetAllBooks).Methods("GET")

	router.HandleFunc("/books/author/{author}", GetBook).Methods("GET")
	router.HandleFunc("/books/title/{title}", GetBookT).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")

	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
