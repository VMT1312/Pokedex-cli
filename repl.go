package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	var err error

	commands := map[string]cliCommand{
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
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println("Error reading input, terminating.")
			break
		}

		text := scanner.Text()
		words := cleanInput(text)

		command, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err = command.callback()
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
