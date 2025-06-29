package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println("Error reading input, terminating.")
			break
		}

		text := scanner.Text()
		words := cleanInput(text)

		fmt.Printf("Your command was: %s\n", words[0])
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
