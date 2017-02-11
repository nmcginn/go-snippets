package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadJson() string {
	var j map[string]interface{}
	data, err := ioutil.ReadFile("data/example.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(data, &j); err != nil {
		log.Fatal(err)
	}
	return j["item"].(string)
}

