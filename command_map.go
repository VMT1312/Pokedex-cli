package main

import (
	"errors"
	"fmt"
)

func commandMapf(s []string, c *config) error {
	res, err := c.pokeClient.ListLocations(c.Next, c.cache)
	if err != nil {
		return err
	}

	c.Next = res.Next
	c.Previous = res.Previous

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(s []string, c *config) error {
	if c.Previous == nil {
		return errors.New("you're on the first page")
	}

	res, err := c.pokeClient.ListLocations(c.Next, c.cache)
	if err != nil {
		return err
	}

	c.Next = res.Next
	c.Previous = res.Previous

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	return nil
}
