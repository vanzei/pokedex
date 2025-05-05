package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/vanzei/pokedex/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	// Create a context
	ctx := context.Background()

	// Generate a cache key based on the URL
	cacheKey := "default_locations"
	if cfg.nextLocationsURL != nil {
		cacheKey = *cfg.nextLocationsURL
	}

	// Check if the response is in the cache
	cachedData, found := cfg.cache.Get(cacheKey)
	if found {
		cfg.apiCallCounter.WithLabelValues("GET", "cacheHit").Inc()
		// Parse the cached data
		var locationResp pokeapi.RespShallowLocations
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return fmt.Errorf("failed to parse cached data: %v", err)
		}

		// Update the next and previous URLs
		cfg.nextLocationsURL = locationResp.Next
		cfg.prevLocationsURL = locationResp.Previous

		// Print the locations
		for _, loc := range locationResp.Results {
			fmt.Println(loc.Name)
		}
		return nil
	}

	// If not in cache, make the API call
	locationsResp, err := cfg.pokeapiClient.ListLocations(ctx, cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// Increment the API call counter
	cfg.apiCallCounter.WithLabelValues("GET", "mapHit").Inc()

	// Cache the response
	cachedData, err = json.Marshal(locationsResp)
	if err != nil {
		return fmt.Errorf("failed to cache response: %v", err)
	}
	cfg.cache.Add(cacheKey, cachedData)

	// Update the next and previous URLs
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// Print the locations
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

	// Generate a cache key based on the URL
	cacheKey := *cfg.prevLocationsURL

	// Check if the response is in the cache
	cachedData, found := cfg.cache.Get(cacheKey)
	if found {
		cfg.apiCallCounter.WithLabelValues("GET", "cacheHit").Inc()
		// Parse the cached data
		var locationResp pokeapi.RespShallowLocations
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return fmt.Errorf("failed to parse cached data: %v", err)
		}

		// Update the next and previous URLs
		cfg.nextLocationsURL = locationResp.Next
		cfg.prevLocationsURL = locationResp.Previous

		// Print the locations
		for _, loc := range locationResp.Results {
			fmt.Println(loc.Name)
		}
		return nil
	}

	// If not in cache, make the API call
	locationResp, err := cfg.pokeapiClient.ListLocations(ctx, cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// Increment the API call counter
	cfg.apiCallCounter.WithLabelValues("GET", "mapbHit").Inc()

	// Cache the response
	cachedData, err = json.Marshal(locationResp)
	if err != nil {
		return fmt.Errorf("failed to cache response: %v", err)
	}
	cfg.cache.Add(cacheKey, cachedData)

	// Update the next and previous URLs
	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	// Print the locations
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
