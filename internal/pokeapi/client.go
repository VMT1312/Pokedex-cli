package pokeapi

import (
	"net/http"
	"time"
)

// Client is a custom HTTP client for the PokeAPI.
type Client struct {
	httpClient http.Client
}

// NewClient creates a new PokeAPI client with a custom timeout.
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
