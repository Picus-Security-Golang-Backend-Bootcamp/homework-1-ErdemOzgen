package jsonops

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func OpenJsonFile(f string) string {
	jsonFile, err := os.Open(f)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened users.json")
	var buf bytes.Buffer
	io.Copy(&buf, jsonFile)
	asString := buf.String()
	//fmt.Println(asString)
	return asString
}
