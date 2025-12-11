package main

import (
	"encoding/json"
	"fmt"
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
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Result   []Link `json:"results"`
}

func commandMap(cfg *config, params []string) error {
	decodedJson := getMapJson(cfg.nextMapUrl, cfg)

	for i := 0; i < len(decodedJson.Result); i++ {
		fmt.Printf("%s\n", decodedJson.Result[i].Name)
	}

	if decodedJson.Next == nil {
		cfg.previousMapUrl = decodedJson.Previous
		return nil
	}
	cfg.nextMapUrl = decodedJson.Next
	cfg.previousMapUrl = decodedJson.Previous
	return nil
}


func commandMapb(cfg *config, params []string) error {
	if cfg.previousMapUrl == nil {
		return fmt.Errorf("youre on the first page")
	}

	decodedJson := getMapJson(cfg.previousMapUrl, cfg)
	for i := 0; i < len(decodedJson.Result); i++ {
		fmt.Printf("%s\n", decodedJson.Result[i].Name)
	}

	if decodedJson.Next == nil {
		cfg.previousMapUrl = decodedJson.Previous
		return nil
	}
	cfg.nextMapUrl = decodedJson.Next
	cfg.previousMapUrl = decodedJson.Previous
	return nil
}

func getMapJson(url *string, cfg *config) ApiResponse {
	if data, inDb := cfg.pokecache.Get(*url); inDb {
		var decodedJson ApiResponse
		json.Unmarshal(data, &decodedJson)
		return decodedJson
	}

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	cfg.pokecache.Add(*url, body)
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
