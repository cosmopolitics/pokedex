package main

import (
	"encoding/json"
	"fmt"
	"github.com/cosmopolitics/cache"
	"io"
	"log"
	"net/http"
)

type Link struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ApiResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Result   []Link `json:"results"`
}

func commandMap(url *string, pokeCache *cache.Cache) string {
	decodedJson := getJson(url, pokeCache)

	for i := 0; i < len(decodedJson.Result); i++ {
		fmt.Printf("%s\n", decodedJson.Result[i].Name)
	}

	if decodedJson.Next == "" {
		return decodedJson.Previous
	}
	*url = decodedJson.Next

	return decodedJson.Previous
}

func getJson(url *string, pokeCache *cache.Cache) ApiResponse {
	if data, inDb := pokeCache.Get(*url); inDb {
		var decodedJson ApiResponse
		json.Unmarshal(data, &decodedJson)
		return decodedJson
	}

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	pokeCache.Add(*url, body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	var decodedjson ApiResponse
	json.Unmarshal(body, &decodedjson)

	return decodedjson
}
