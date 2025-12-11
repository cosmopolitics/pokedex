package main

import (
	"fmt"
	"os"
	"strings"
)

type command struct {
	name string
	description string
	callback func(*config) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands(cfg *config) map[string]command {
	return map[string]command{
		"exit": {
			name: "exit",
			description: "exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "prints help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "gets next twenty locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "gets previous page of locations",
			callback: commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	_, err := fmt.Printf("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	commands := getCommands(cfg)
	for _, h := range commands {
		fmt.Printf("%s: %s\n", h.name, h.description)
	}
	return nil
}

