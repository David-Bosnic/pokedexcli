package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/David-Bosnic/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

type Config struct {
	Next string
	Prev string
}

type PokeMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var commands map[string]cliCommand
var pokeCache *internal.Cache

func init() {
	pokeCache = internal.NewCache(5)
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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon world",
			callback:    commandMapB,
		},
	}
}
func main() {
	currentConfig := Config{
		Next: "",
		Prev: "",
	}
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
		if len(cleanText) != 0 {
			_, ok := commands[cleanText[0]]
			if ok {
				commands[cleanText[0]].callback(&currentConfig)
			} else {
				fmt.Println("Unknown Command")
			}

		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	stringers := strings.Fields(text)
	return stringers
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(config *Config) error {
	// TODO: Make the list ordered rather than rng
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}
func commandMap(config *Config) error {
	var url string
	var pokeMap PokeMap
	if config.Next != "" {
		url = config.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	val, ok := pokeCache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 200 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return err
		}
		pokeMap = PokeMap{}
		pokeCache.Add(url, body)
		err = json.Unmarshal(body, &pokeMap)
		if err != nil {
			return err
		}
	} else {
		err := json.Unmarshal(val, &pokeMap)
		if err != nil {
			return err
		}
	}
	for _, val := range pokeMap.Results {
		fmt.Println(val.Name)
	}
	config.Next = pokeMap.Next
	config.Prev = pokeMap.Previous
	return nil
}

func commandMapB(config *Config) error {
	var url string
	var pokeMap PokeMap
	if config.Prev != "" {
		url = config.Prev
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	val, ok := pokeCache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 200 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return err
		}
		pokeMap = PokeMap{}
		pokeCache.Add(url, body)
		err = json.Unmarshal(body, &pokeMap)
		if err != nil {
			return err
		}
	} else {
		err := json.Unmarshal(val, &pokeMap)
		if err != nil {
			return err
		}
	}
	for _, val := range pokeMap.Results {
		fmt.Println(val.Name)
	}
	config.Next = pokeMap.Next
	config.Prev = pokeMap.Previous
	return nil
}
