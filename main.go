package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show commands",
			callback:    commandHelp,
		},
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		cleanText := cleanInput(scanner.Text())
		_, ok := commands[cleanText[0]]
		if ok {
			commands[cleanText[0]].callback()
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	stringers := strings.Fields(text)
	return stringers
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}
