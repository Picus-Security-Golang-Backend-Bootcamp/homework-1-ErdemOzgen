package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
)

func PrintPretty(book interface{}) {
	b, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}

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
func RunServer() {

	cmd := exec.Command("go", "run", "./restapi/restapiServer.go")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("erdem")
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		fmt.Println(sig)
		done <- true
	}()
	err := cmd.Start()

	if err != nil {
		log.Fatal(err)
	}
	openbrowser("http://localhost:8000/books")
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
