package main

import "fmt"

func main() {
	// different ways to declare a map
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#FF1212",
	}

	colors["white"] = "#FFFFFF"

	delete(colors, "white")

	fmt.Println(colors)
}
