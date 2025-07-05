package main

import (
	"time"

	"github.com/VMT1312/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeClient: pokeClient,
	}

	startRepl(config)
}
