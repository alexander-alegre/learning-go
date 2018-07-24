package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	res.Body.Read(bs)

	fmt.Println(string(bs))
}
