package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/VMT1312/pokedexcli/internal/pokeapi"
	"github.com/VMT1312/pokedexcli/internal/pokecache"
)

type config struct {
	Next       *string
	Previous   *string
	pokeClient pokeapi.Client
	cache      *pokecache.Cache
	pokedex    map[string]pokeapi.Pokemon
}

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println("Error reading input, terminating.")
			break
		}

		text := scanner.Text()
		words := cleanInput(text)

		command, ok := getCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(words, c)
			if err != nil {
				fmt.Printf("Error executing command '%s': %v\n", command.name, err)
				continue
			}
		}

	}
}

func cleanInput(text string) []string {
	if text == "" {
		println("Input is empty")
		return []string{}
	}

	lower_case := strings.ToLower(text)

	words := strings.Fields(lower_case)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func([]string, *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"mapf": {
			name:        "mapf",
			description: "Displays the next page of the region in Pokemon",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of the region in Pokemon",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all the Pokemons within a region",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
	}
}
