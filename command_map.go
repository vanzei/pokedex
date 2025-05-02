package main

import (
	"context"
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	// Create a context
	ctx := context.Background()

	// Pass the context to ListLocations
	locationsResp, err := cfg.pokeapiClient.ListLocations(ctx, cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	// Create a context
	ctx := context.Background()

	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// Pass the context to ListLocations
	locationResp, err := cfg.pokeapiClient.ListLocations(ctx, cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
