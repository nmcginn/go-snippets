package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	var y map[string]interface{}
	data, err := ioutil.ReadFile("data/example.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(data, &y); err != nil {
		log.Fatal(err)
	}
	fmt.Println(y["item"].(string))
}

