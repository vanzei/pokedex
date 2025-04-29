package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)



func startRepl() {

    fmt.Print("Pokedex > ")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        input := cleanInput(scanner.Text())

        if len(input) == 0 {
            fmt.Print("Pokedex > ")
            continue
        }

        // Get the command from the input
        command := input[0]

        // Check if the command exists in the commands map
        if cmd, exists := getCommands()[command]; exists {
            // Execute the command's callback function
            if err := cmd.callback(); err != nil {
                fmt.Printf("Error executing command '%s': %v\n", command, err)
            }
        } else {
            // Handle unknown commands
            fmt.Printf("Unknown command: %s\n", command)
        }

        fmt.Print("Pokedex > ")
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error reading input:", err)
        os.Exit(1)
    }
}

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}




type cliCommand struct {
    name        string
    description string
    callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}