package main

import "strings"

func cleanInput(text string) []string {
	if text == "" {
		println("Input is empty")
		return []string{}
	}

	lower_case := strings.ToLower(text)

	words := strings.Fields(lower_case)

	return words
}
