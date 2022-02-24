package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func PrintPretty(book interface{}) {
	b, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}

func Openbrowser(url string) {
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
func RunServer() {

	cmd := exec.Command("go", "run", "./restapi/restapiServer.go")

	err := cmd.Start()

	if err != nil {
		log.Fatal(err)
	}
}
