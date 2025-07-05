package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/VMT1312/pokedexcli/internal/pokecache"
)

// List locations
func (c *Client) ListPokemon(location string, cache *pokecache.Cache) (Pokemons, error) {
	if location == "" {
		return Pokemons{}, errors.New("location cannot be empty")
	}
	url := baseURL + "/location-area/" + location + "/"

	var body []byte

	if cachedBody, found := cache.Get(url); found {
		body = cachedBody
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemons{}, err
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemons{}, err
		}

		cache.Add(url, body)
	}

	var pokemons Pokemons
	if err := json.Unmarshal(body, &pokemons); err != nil {
		return Pokemons{}, err
	}
	return pokemons, nil
}
