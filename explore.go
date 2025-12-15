package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


func commandExplore(cfg *config, params []string) error {
	if len(params) < 2 {
		fmt.Println("Usage:\n\texplore <region/location name or id>")
	}
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
