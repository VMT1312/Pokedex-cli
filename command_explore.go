package main

import "fmt"

func commandExplore(location []string, c *config) error {
	res, err := c.pokeClient.ListPokemon(location[1], c.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")

	for _, encouter := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", encouter.Pokemon.Name)
	}

	return nil
}
