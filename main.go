package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	stringers := strings.Fields(text)
	return stringers
}
