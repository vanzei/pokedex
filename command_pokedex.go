package main

import (
    "fmt"
    "sort"
)

func commandPokedex(cfg *config) error {
    if len(cfg.pokemonList) == 0 {
        fmt.Println("Your Pokedex is empty. Catch some Pokemon!")
        return nil
    }

    fmt.Println("Your Pokedex:")
    fmt.Println("-------------")

    // Get a sorted list of Pokemon names for consistent output
    var names []string
    for name := range cfg.pokemonList {
        names = append(names, name)
    }
    sort.Strings(names)

    // Print each Pokemon
    for _, name := range names {
        fmt.Printf("  - %s\n", name)
    }
    
    fmt.Printf("\nYou have caught %d Pokemon!\n", len(cfg.pokemonList))
	cfg.apiCallCounter.WithLabelValues("GET", "pokedex").Inc()
    return nil
}