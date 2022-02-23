package restapi

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/model"
	"net/http"

	"github.com/gorilla/mux"
)

type Book model.Book

var books []Book

func GetBookByAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, b := range books {

		if b.Author == params["author"] {
			fmt.Println(b.Author)
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	json.NewEncoder(w).Encode("Book not found-> Searched for author")
}

func GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, b := range books {

		if b.Author == params["title"] {
			fmt.Println(b.Author)
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	json.NewEncoder(w).Encode("Book not found-> Searched for title")
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}
