package main

import (
	"errors"
	"fmt"
)

func commandInspect(pokemon []string, c *config) error {
	if len(pokemon) < 2 {
		return errors.New("please provide a pokemon name")
	}

	pokemon_info, ok := c.pokedex[pokemon[1]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon_info.Name)

	fmt.Printf("Height: %d\n", pokemon_info.Height)

	fmt.Printf("Weight: %d\n", pokemon_info.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon_info.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon_info.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
