package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/VMT1312/pokedexcli/internal/pokecache"
)

// List locations
func (c *Client) GetPokemon(pokemon string, cache *pokecache.Cache) (Pokemon, error) {
	if pokemon == "" {
		return Pokemon{}, errors.New("location cannot be empty")
	}
	url := baseURL + "/pokemon/" + pokemon + "/"

	var body []byte

	if cachedBody, found := cache.Get(url); found {
		body = cachedBody
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		cache.Add(url, body)
	}

	var pokemon_info Pokemon
	if err := json.Unmarshal(body, &pokemon_info); err != nil {
		return Pokemon{}, err
	}
	return pokemon_info, nil
}
