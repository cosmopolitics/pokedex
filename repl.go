package main

import (
	"fmt"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	callback    func(cfg *config, params []string) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands(cfg *config) map[string]command {
	return map[string]command{
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "prints help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "gets next twenty locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "gets previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "lists all pokemon for given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "throws a pokeball at a specified pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "get stats of pokemon in your pokedex (caught)",
			callback:    commandInspect,
		},
	}
}

func commandExit(cfg *config, params []string) error {
	_, err := fmt.Printf("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, params []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	commands := getCommands(cfg)
	for _, h := range commands {
		fmt.Printf("%s: %s\n", h.name, h.description)
	}
	return nil
}
