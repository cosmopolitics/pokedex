package main

import (
	"net/http"
	"io"
	"encoding/json"
	"fmt"
	"log"
)

type Link struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type JsonList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Result   []Link `json:"results"`
}

func commandMap(url *string) string {
	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	var decodedjson JsonList
	json.Unmarshal(body, &decodedjson)

	for i := 0;i < len(decodedjson.Result); i++ {
		fmt.Printf("%s\n", decodedjson.Result[i].Name)
	}

	if decodedjson.Next == "" {
		return decodedjson.Previous
	}
	*url = decodedjson.Next

	return decodedjson.Previous
}
