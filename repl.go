package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cosmopolitics/pokecache"
)

type command struct {
	name string
	description string
	callback func() error
}

type config struct {
	previousMapUrl *string
	nextMapUrl *string
	pokecache pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration)

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
			name: "mapb",
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
	for _, h := range commands {
		fmt.Printf("%s: %s\n", h.name, h.description)
	}
	return nil
}

