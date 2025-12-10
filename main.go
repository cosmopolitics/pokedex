package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/cosmopolitics/cache"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	mapUrl := "https://pokeapi.co/api/v2/location-area"
	previousMapUrl := ""

	fmt.Println("Welcome to the Pokedex!")
	for {
		// Prompt
		green := "\033[32m"
		reset := "\033[0m"
		fmt.Print(green + "Pokedex: " + reset)

		reader.Scan()
		cleanText := cleanInput(reader.Text())
		if len(cleanText) == 0 {
			continue
		}

		// Do command
		commands := getCommands()

		if cleanText[0] == "map" {
			previousMapUrl = commandMap(&mapUrl)
		} else if cleanText[0] == "mapb" {
			if previousMapUrl == "" {
				fmt.Printf("no previous map page\n")
				continue
			}
			previousMapUrl = commandMap(&previousMapUrl)
		} else if commands[cleanText[0]].callback != nil {
			commands[cleanText[0]].callback()
		}
	}
}


