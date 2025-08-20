package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("counter keep tracks of the count", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)

	})

	t.Run("counter runs safely concurrently", func(t *testing.T) {
		repetitions := 1000
		counter := Counter{}
		var wg sync.WaitGroup
		wg.Add(repetitions)
		for i := 0; i < repetitions; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, &counter, repetitions)
	})
}

func assertCounter(t testing.TB, counter *Counter, expected int) {
	if counter.Value() != expected {
		t.Errorf("expected counter to be called %d times, but got %v", expected, counter.Value())
	}
}
