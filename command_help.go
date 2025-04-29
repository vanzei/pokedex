package main

import (
	"fmt"
)


func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for key, option := range getCommands() {
        fmt.Printf("  %s: %s\n", key, option.description)
    }
    return nil
}