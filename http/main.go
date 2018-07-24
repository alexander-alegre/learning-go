package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	lw := logWriter{}
	// this line below does the same thing as the code commented out
	io.Copy(lw, res.Body)

	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote %v bytes \n", len(bs))

	return len(bs), nil
}
