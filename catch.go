package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func commandCatch(cfg *config, params []string) error {
	if len(params) < 2 {
		fmt.Println("Usage:\n\tcatch <pokemon name or id>")
	}

	baseUrl := "https://pokeapi.co/api/v2/"
	body, _ := makeApiGet(baseUrl + "pokemon/" + params[1], cfg)
	var pokemon Pokemon_endpoint
	err := json.Unmarshal(body, &pokemon)

	if err != nil {
		body, err := makeApiGet(baseUrl + "pokemon-species/" + params[1], cfg)
		if err != nil {
			fmt.Printf("couldnt find the pokemon")
			return nil
		}
		var species_data Pokemon_Species
		json.Unmarshal(body, &species_data)

		fmt.Println("there are multiple pokedex entries please use one of the following ids")
		for _, p := range species_data.PokedexNumbers {
			fmt.Println(p.Pokedex.Name)
			fmt.Println(p.EntryNumber)
		}
		return nil
	} else {
		fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
		res := rand.Intn(pokemon.BaseExperience)

		if res < 40 {
			fmt.Println("you caught it :3")
			cfg.their_pokedex[pokemon.Name] = pokemon
			fmt.Printf("%s was added to your pokedex\n", pokemon.Name)
			return nil
		}
		fmt.Println("you failed :(")
	}
	return nil
}

func makeApiGet(url string, cfg *config) ([]byte, error) {
	if data, cached := cfg.pokecache.Get(url); cached {
		return data, nil
	}
	res, err := http.Get(url)
	if res.StatusCode == 404 {
		return nil, errors.New("404")
	}
	if res.StatusCode > 299 || res.StatusCode < 200 {
		log.Fatalf("%d, not successful", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, errors.New("failed to read response")
	}

	cfg.pokecache.Add(url, body)
	return body, nil
}

