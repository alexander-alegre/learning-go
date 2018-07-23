package main

import (
	"fmt"
)

func main() {
	// different ways to declare a map
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#FF1212",
		"black": "F000000",
	}

	// adding a new value
	colors["white"] = "#FFFFFF"

	// deleting a value
	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("color: %v hex: %v \n", color, hex)
	}
}
