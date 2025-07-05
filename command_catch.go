package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(pokemon []string, c *config) error {
	res, err := c.pokeClient.GetPokemon(pokemon[1], c.cache)
	if err != nil {
		return err
	}

	catch_chance := rand.Intn(100)
	base_chance := 100
	var chance int
	if base_chance-res.BaseExperience < 0 {
		chance = 5
	} else if base_chance-res.BaseExperience > 100 {
		chance = 95
	} else {
		chance = base_chance - res.BaseExperience
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon[1])

	if chance > catch_chance {
		fmt.Printf("%s was caught!\n", res.Name)
		c.pokedex[pokemon[1]] = res
	} else {
		fmt.Printf("%s escaped!\n", res.Name)
	}

	return nil
}
