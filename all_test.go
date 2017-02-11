package main

import (
	"testing"
)

func TestJson(t *testing.T) {
	str := ReadJson()
	if str != "somedata" {
		t.Error("JSON did not read correctly.")
	}
}

func TestYaml(t *testing.T) {
	str := ReadYaml()
	if str != "somedata" {
		t.Error("YAML did not read correctly.")
	}
}
