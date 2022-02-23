package main

import (
	"encoding/json"
	"fmt"
	"homework-1-ErdemOzgen/jsonops"
	"homework-1-ErdemOzgen/model"
	"homework-1-ErdemOzgen/search"
	"log"
	"os/exec"
	"runtime"
)

type Book model.Book // Questtion []search.Book is good implementation???
var books []Book

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
func runServer() {

	cmd := exec.Command("go", "run", "./restapi/restapiServer.go")

	err := cmd.Start()

	if err != nil {
		log.Fatal(err)
	}
}
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
	runServer()
	openbrowser("http://localhost:8080")
}
