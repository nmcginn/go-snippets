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

func TestReverse(t *testing.T) {
	test_string := "my string"
	reversed := ReverseString(test_string)
	if reversed != "gnirts ym" {
		t.Error("String was not reversed correctly.")
	}
}

func TestYaml(t *testing.T) {
	str := ReadYaml()
	if str != "somedata" {
		t.Error("YAML did not read correctly.")
	}
}
