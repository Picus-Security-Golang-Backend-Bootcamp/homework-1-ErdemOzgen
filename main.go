package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func OpenJsonFile() {
	jsonFile, err := os.Open("test.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened users.json")
	var buf bytes.Buffer
	io.Copy(&buf, jsonFile)
	asString := buf.String()
	fmt.Println(asString)

}

func main() {
	fmt.Println("Staring Program")
	OpenJsonFile()
}
