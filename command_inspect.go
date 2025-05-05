package main

import (
    "fmt"
    "strings"
)

func commandInspect(cfg *config) error {
    // Check for arguments
    if len(cfg.commandArgs) == 0 {
        return fmt.Errorf("missing pokemon name")
    }
    pokemonName := strings.ToLower(cfg.commandArgs[0])
    
    // Check if the pokemon is caught
    pokemon, found := cfg.pokemonList[pokemonName]
    if !found {
        fmt.Println("You have not caught that pokemon")
        return nil
    }
    
    // Display pokemon details
    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)
    
    // Display stats
    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
    }
    
    // Display types
    fmt.Println("Types:")
    for _, typeInfo := range pokemon.Types {
        fmt.Printf("  - %s\n", typeInfo.Type.Name)
    }
    cfg.apiCallCounter.WithLabelValues("GET", "inspect").Inc()
    return nil
}