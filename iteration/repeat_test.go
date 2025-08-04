package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat returns the correct value", func(t *testing.T) {
		got := Repeat("a", 2)
		expected := "aa"
		assertionHelper(t, got, expected)
	})

	t.Run("Repeat iterates as many times as specified", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"
		assertionHelper(t, got, expected)
	})
}

func assertionHelper(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	got := Repeat("sanjay", 2)
	fmt.Println(got)
	// Output: sanjaysanjay
}