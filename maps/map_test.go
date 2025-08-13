package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}

	t.Run("known word", func(t *testing.T){
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("unknown word", func(t *testing.T){
		_, err := dictionary.Search("sanjay")
		assertError(t, err, ErrUnkownWord)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "sanjay"
		definition := "an awesome software dev"
		dict.Add(word, definition)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T){
		word := "sanjay"
		definition := "an awesome software dev"
		dict := Dictionary{word: definition}
		err := dict.Add(word, "not awesome")
		assertError(t, err, ErrExistingWord)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T){
		word := "go"
		def := "gopher"
		dict := Dictionary{word: def}
		newDef := "an amazing language"
		err := dict.Update(word, newDef)
		assertNoError(t, err)
		assertDefinition(t, dict, word, newDef)
	})

	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update("sanjay", "a nice guy")
		assertError(t, err, ErrUnkownWord)
	})
}

func assertDefinition(t testing.TB, d Dictionary, word, expected string) {
	t.Helper()
	got, err := d.Search(word)
	assertNoError(t, err)
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func assertError(t testing.TB, err, expected error) {
	if err == nil {
		t.Fatal("expected an error but got no error")
	}
	if err != expected {
		t.Errorf("expected error %q but got %q", err, expected)
	}
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Fatal("expected no error but got error")
	}
}
