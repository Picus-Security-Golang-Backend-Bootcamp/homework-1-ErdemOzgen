package utils

import (
	"encoding/json"
	"fmt"
)

func PrintPretty(book interface{}) {
	b, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}
