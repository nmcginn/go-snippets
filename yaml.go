package main

import (
		"fmt"
		"gopkg.in/yaml.v2"
		"io/ioutil"
		"log"
)

type yamlfile struct {
	item string   `yaml:"item"`
	list []string `yaml:"list"`
}

func main() {
	y := yamlfile{}
	data, err := ioutil.ReadFile("data/example.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(data, &y); err != nil {
		log.Fatal(err)
	}
	fmt.Println(y.item)
	fmt.Println(y.list)
}

