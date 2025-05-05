package main

import (
	"fmt"
)


func commandHelp(cfg *config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for key, option := range getCommands() {
        fmt.Printf("  %s: %s\n", key, option.description)
    }
    cfg.apiCallCounter.WithLabelValues("GET", "help").Inc()
    return nil
}