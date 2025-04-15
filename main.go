package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		cleanText := cleanInput(scanner.Text())
		if len(cleanText) > 0 {
			fmt.Println("Your command was:", cleanText[0])
		} else {
			fmt.Println()
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	stringers := strings.Fields(text)
	return stringers
}
