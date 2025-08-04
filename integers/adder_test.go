package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2,2)
	want := 4
	if got != want {
		t.Errorf("got '%d', want '%d'", got, want)
	}
}

func ExampleAdd() {
	got := Add(3,3)
	fmt.Println(got)
	// Output: 6
}