package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://golang.org",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%v might be down \n", link)
		c <- "Might be down I think"
		return
	}

	fmt.Printf("%v is up! \n", link)
	c <- "Yep, it's up"
}
