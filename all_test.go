package main

import (
	"testing"
)

func TestYaml(t *testing.T) {
	str := ReadYaml()
	if str != "somedata" {
		t.Error("YAML did not read correctly.")
	}
}
