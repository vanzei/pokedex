package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vanzei/pokedex/internal/pokeapi"
    "github.com/vanzei/pokedex/internal/pokecache" // Use the full package name
    "github.com/prometheus/client_golang/prometheus"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokemonList      map[string]pokeapi.PokemonResponse // Changed from []string
	apiCallCounter   *prometheus.CounterVec
	cache            *pokecache.Cache
	commandArgs      []string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:] // Get any arguments after the command

		command, exists := getCommands()[commandName]
		if exists {
			// Store arguments in the config
			cfg.commandArgs = args
			
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Get the next page of locations",
            callback:    commandMapf,
        },
        "mapb": {
            name:        "mapb",
            description: "Get the previous page of locations",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore",
            description: "Explore a location area to find Pokémon (usage: explore <location-area-name>)",
            callback:    commandExplore,
        },
		 "inspect": {
            name:        "inspect",
            description: "View details of a caught pokemon (usage: inspect <pokemon-name>)",
            callback:    commandInspect,
        },
		"catch": {
            name:        "catch",
            description: "Attempt to catch a pokemon (usage: catch <pokemon-name>)",
            callback:    commandCatch,
        },
		"pokedex": {
            name:        "pokedex",
			description: "View the list of caught Pokémon",
            callback:    commandPokedex,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
    }
}
