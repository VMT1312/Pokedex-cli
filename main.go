package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
