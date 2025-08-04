package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum returns the sum of all the elements in the collection", func(t *testing.T) {
		input := []int{1,2,3,4,5}
		got := Sum(input)
		expected := 15
		if got != expected {
			t.Errorf("Expected '%d' but got '%d' input %v", expected, got, input)
		}
	})

	t.Run("Sum can accept a dynamic number of elements", func(t *testing.T) {
		input := []int{1,2,3}
		got := Sum(input)
		expected := 6
		if got != expected {
			t.Errorf("Expected '%d' but got '%d' input '%v'", expected, got, input)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	b.Run("benchmark 1", func(b *testing.B) {
		input := []int{1, 2}
		for range b.N {
			Sum(input)
		}
	})
	
	b.Run("benchmark 2", func(b *testing.B) {
		input := []int{1, 2,3,4,5}
		for range b.N {
			Sum(input)
		}
	})
}