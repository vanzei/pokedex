package main

import (
    "context"
    "encoding/json"
    "fmt"
    "strings"

    "github.com/vanzei/pokedex/internal/pokeapi"
)

func commandExplore(cfg *config) error {
    // Get the location area name from the arguments
    if len(cfg.commandArgs) == 0 {
        return fmt.Errorf("missing location area name")
    }
    locationAreaName := strings.ToLower(cfg.commandArgs[0])

    // Create a context
    ctx := context.Background()

    // Generate a cache key based on the location area name
    cacheKey := fmt.Sprintf("location-area:%s", locationAreaName)

    // Check if the response is in the cache
    cachedData, found := cfg.cache.Get(cacheKey)
    if found {
        // Increment cache hit counter
        cfg.apiCallCounter.WithLabelValues("GET", "exploreCache").Inc()
        
        // Parse the cached data
        var locationArea pokeapi.LocationAreaResponse
        err := json.Unmarshal(cachedData, &locationArea)
        if err != nil {
            return fmt.Errorf("failed to parse cached data: %v", err)
        }

        // Display the location area and its Pokémon
        fmt.Printf("Exploring %s...\n", locationArea.Name)
        fmt.Println("Found Pokémon:")
        for _, encounter := range locationArea.PokemonEncounters {
            fmt.Printf("- %s\n", encounter.Pokemon.Name)
        }
        cfg.apiCallCounter.WithLabelValues("GET", "pokedex").Inc()
        return nil
    }

    // If not in cache, make the API call
    locationArea, err := cfg.pokeapiClient.GetLocationArea(ctx, locationAreaName)
    if err != nil {
        return fmt.Errorf("failed to explore location area: %v", err)
    }

    // Increment the API call counter
    cfg.apiCallCounter.WithLabelValues("GET", "explore").Inc()

    // Cache the response
    cachedData, err = json.Marshal(locationArea)
    if err != nil {
        return fmt.Errorf("failed to cache response: %v", err)
    }
    cfg.cache.Add(cacheKey, cachedData)

    // Display the location area and its Pokémon
    fmt.Printf("Exploring %s...\n", locationArea.Name)
    fmt.Println("Found Pokémon:")
    for _, encounter := range locationArea.PokemonEncounters {
        fmt.Printf("- %s\n", encounter.Pokemon.Name)
    }
    return nil
}