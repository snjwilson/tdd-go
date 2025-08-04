package main

import "testing"

func TestHello(t *testing.T){
	t.Run("say hello to a person", func(t *testing.T) {
		got:=hello("Sanjay", "")
		want:="Hello Sanjay!"
	assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to world when there is no person", func(t *testing.T) {
		got := hello("", "English")
		want := "Hello world!"

	assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got:=hello("Sanjay", "Spanish")
		want:="Hola Sanjay!"
		assertCorrectMessage(t,got,want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got:=hello("Sanjay", "French")
		want:="Bonjour Sanjay!"
		assertCorrectMessage(t,got,want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

