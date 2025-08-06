package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum returns the sum of all the elements in the collection", func(t *testing.T) {
		input := []int{1,2,3,4,5}
		got := Sum(input)
		expected := 15
		assertionHelperSum(t,got, expected, input)
	})

	t.Run("Sum can accept a dynamic number of elements", func(t *testing.T) {
		input := []int{1,2,3}
		got := Sum(input)
		expected := 6
		assertionHelperSum(t,got, expected, input)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("SumAll should return the correct sum of all the slices", func(t *testing.T) {
		got := SumAll([]int{1,2,3}, []int{1,2,3,4,5})
		expected := []int{6,15}
		assertionHelperSumAll(t, got, expected)
	})
}

func TestSumAllTrail(t *testing.T) {
	t.Run("SumAll should return the correct sum of all the trail of the slices", func(t *testing.T) {
		got := SumAllTail([]int{1,2,3}, []int{1,2,3,4,5})
		expected := []int{5,14}
		assertionHelperSumAll(t, got, expected)
	})
}


func assertionHelperSum(t testing.TB, got, expected int, input []int) {
	t.Helper()
	if got != expected {
		t.Errorf("Expected '%d' but got '%d' input '%v'", expected, got, input)
	}
}

func assertionHelperSumAll(t testing.TB, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
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