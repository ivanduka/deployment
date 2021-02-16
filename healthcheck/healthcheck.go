package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := 3333
	resp, err := http.Get(fmt.Sprintf("http://localhost:%v", port))
	if err != nil || (resp == nil || resp.StatusCode != http.StatusOK) {
		fmt.Println("ERROR")
		os.Exit(1)
	}
	fmt.Println("OK")
	os.Exit(0)
}