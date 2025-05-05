package main

import (
    "context"
    "encoding/json"
    "fmt"
    "math"
    "math/rand"
    "strings"
    "time"

    "github.com/vanzei/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config) error {
    // Get the pokemon name from the arguments
    if len(cfg.commandArgs) == 0 {
        return fmt.Errorf("missing pokemon name")
    }
    pokemonName := strings.ToLower(cfg.commandArgs[0])

    // Create a context
    ctx := context.Background()

    // Generate a cache key based on the pokemon name
    cacheKey := fmt.Sprintf("pokemon:%s", pokemonName)

    // Check if the response is in the cache
    cachedData, found := cfg.cache.Get(cacheKey)
    var pokemon pokeapi.PokemonResponse
    
    if found {
        // Parse the cached data
        err := json.Unmarshal(cachedData, &pokemon)
        if err != nil {
            return fmt.Errorf("failed to parse cached data: %v", err)
        }
    } else {
        // If not in cache, make the API call
        var err error
        pokemon, err = cfg.pokeapiClient.GetPokemon(ctx, pokemonName)
        if err != nil {
            return fmt.Errorf("failed to get pokemon: %v", err)
        }

        // Increment the API call counter
        cfg.apiCallCounter.WithLabelValues("GET", "catch").Inc()

        // Cache the response
        cachedData, err = json.Marshal(pokemon)
        if err != nil {
            return fmt.Errorf("failed to cache response: %v", err)
        }
        cfg.cache.Add(cacheKey, cachedData)
    }

    // Check if already caught
    if _, alreadyCaught := cfg.pokemonList[pokemon.Name]; alreadyCaught {
        fmt.Printf("You already have %s!\n", pokemon.Name)
        return nil
    }

    // Attempt to catch the pokemon
    fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
    
    // Create a short suspense
    time.Sleep(1 * time.Second)

    // Try to catch the pokemon
    caught := attemptCatch(pokemon.BaseExperience)

    if caught {
        fmt.Printf("Caught %s!\n", pokemon.Name)
        // Store full pokemon data instead of just the name
        cfg.pokemonList[pokemon.Name] = pokemon
    } else {
        fmt.Printf("%s broke free!\n", pokemon.Name)
    }
    
    return nil
}

// attemptCatch determines whether the pokemon is caught based on its base experience
func attemptCatch(baseExperience int) bool {
    // Higher base experience means harder to catch
    // Calculate catch difficulty between 0 and 1
    // Higher value means harder to catch
    difficulty := math.Min(float64(baseExperience)/400.0, 0.9)
    
    // Random number between 0 and 1
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    randomValue := r.Float64()
    
    // If random value is greater than difficulty, catch is successful
    return randomValue > difficulty
}