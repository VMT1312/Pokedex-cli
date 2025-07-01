package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type locationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	res, err := http.Get(*c.Next)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}

	var location_area locationArea
	if err := json.Unmarshal(body, &location_area); err != nil {
		log.Fatal(err)
	}

	c.Next = location_area.Next
	c.Previous = location_area.Previous

	for _, area := range location_area.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(*c.Previous)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}

	var location_area locationArea
	if err := json.Unmarshal(body, &location_area); err != nil {
		log.Fatal(err)
	}

	c.Next = location_area.Next
	c.Previous = location_area.Previous

	for _, area := range location_area.Results {
		fmt.Println(area.Name)
	}

	return nil
}
