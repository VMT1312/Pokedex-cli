package main

import (
	"time"

	"github.com/VMT1312/pokedexcli/internal/pokeapi"
	"github.com/VMT1312/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeClient: pokeClient,
		cache:      pokecache.NewCache(5 * time.Second),
		pokedex:    make(map[string]pokeapi.Pokemon),
	}

	startRepl(config)
}
