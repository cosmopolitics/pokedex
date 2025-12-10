package main

import (
	"fmt"
	"os"
	"strings"
)

type command struct {
	name string
	description string
	callback func() error
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]command {
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
			callback: nil,
			// called manually to implement page memory
		},
		"mapb": {
			name: "map",
			description: "gets previous page of locations",
			callback: nil,
			// called manually to implement page memory
		},
	}
}

func commandExit() error {
	_, err := fmt.Printf("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	commands := getCommands()
	for k, _ := range commands {
		fmt.Printf("%s: %s\n", k, commands[k].description)
	}
	return nil
}

