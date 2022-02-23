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
		//fmt.Println(params["author"])
		if b.Author == params["author"] {
			fmt.Println(b.Author)
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	json.NewEncoder(w).Encode("Book not found")
}
