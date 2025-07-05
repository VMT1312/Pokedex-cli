package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/VMT1312/pokedexcli/internal/pokecache"
)

// List locations
func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (locationArea, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	var body []byte

	if cachedBody, found := cache.Get(url); found {
		body = cachedBody
	} else {
		res, err := http.Get(url)
		if err != nil {
			return locationArea{}, err
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return locationArea{}, err
		}

		cache.Add(url, body)
	}

	var location_area locationArea
	if err := json.Unmarshal(body, &location_area); err != nil {
		return locationArea{}, err
	}
	return location_area, nil
}
