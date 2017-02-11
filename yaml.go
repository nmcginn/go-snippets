package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ReadYaml() string {
	var y map[string]interface{}
	data, err := ioutil.ReadFile("data/example.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(data, &y); err != nil {
		log.Fatal(err)
	}
	return y["item"].(string)
}

