package main

import (
	"fmt"
)

func commandInspect(cfg *config, params []string) error {
	if len(params) < 2 {
		fmt.Println("Usage:\n\tinspect <pokemon you have in your pokedex>")
	}
	if pokemon, inPd := cfg.their_pokedex[params[1]]; inPd {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, Type := range pokemon.Types {
			fmt.Printf("  - %s\n", Type.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
