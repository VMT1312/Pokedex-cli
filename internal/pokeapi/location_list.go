package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// List locations
func (c *Client) ListLocations(pageURL *string) (locationArea, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return locationArea{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationArea{}, err
	}

	var location_area locationArea
	if err := json.Unmarshal(body, &location_area); err != nil {
		return locationArea{}, err
	}

	return location_area, nil
}
