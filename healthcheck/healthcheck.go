package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := 3333
	resp, err := http.Get(fmt.Sprintf("http://localhost:%v", port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		fmt.Println("ERROR", resp.StatusCode)
		os.Exit(1)
	}

	fmt.Println("OK", resp.StatusCode)
	os.Exit(0)
}
