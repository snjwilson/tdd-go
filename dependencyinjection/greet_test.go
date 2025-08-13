package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sanjay")
	got := buffer.String()
	expected := "Hello Sanjay!"
	if got != expected {
		t.Errorf("got %q but expected %q", got, expected)
	}
}