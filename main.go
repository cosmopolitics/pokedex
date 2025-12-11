package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/cosmopolitics/pokecache"
)

type config struct {
	previousMapUrl *string
	nextMapUrl *string
	pokecache pokecache.Cache
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	mapUrl := "https://pokeapi.co/api/v2/location-area"

	cfg := &config{
		pokecache: pokecache.NewCache(20 * time.Second),
		nextMapUrl: &mapUrl,
	}

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
		commands := getCommands(cfg)
		if cmd, exist := commands[cleanText[0]]; exist {
			err := cmd.callback(cfg, cleanText)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("%s doesnt exist, 'help' for usage\n", cleanText[0])
		}
	}
}
