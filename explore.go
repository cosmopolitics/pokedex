package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location_endpoint struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, params []string) error {
	for i, v := range params {
		if i == 0 {
			continue
		}
		fmt.Printf("exploring %s...\n", params[i])
		location_endpoint := getLocationJson(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", v), cfg)
		for _, v := range location_endpoint.PokemonEncounters {
			fmt.Println(v.Pokemon.Name)
		}
	}
	return nil
}

func getLocationJson(url string, cfg *config) Location_endpoint {
	if data, inDb := cfg.pokecache.Get(url); inDb {
		var decodedJson Location_endpoint
		json.Unmarshal(data, &decodedJson)
		return decodedJson
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	cfg.pokecache.Add(url, body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	var decodedjson Location_endpoint
	err = json.Unmarshal(body, &decodedjson)
	if err != nil {
		log.Fatalf("failed to unmarshal data: %v", err)
	}

	return decodedjson
}
