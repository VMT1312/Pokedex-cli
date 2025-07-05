package main

import (
	"errors"
	"fmt"
	"maps"
)

func commandPokedex(s []string, c *config) error {
	if len(s) < 2 {
		return errors.New("please provide a pokemon name")
	}

	fmt.Println("Your Pokedex:")

	for k := range maps.Keys(c.pokedex) {
		fmt.Printf("  - %s\n", k)
	}

	return nil
}
