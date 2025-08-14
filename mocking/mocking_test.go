package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}
	Countdown(buffer, sleeper)
	got := buffer.String()
	expected := `3
2
1
Go!`
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}

	if sleeper.Calls != 3 {
		t.Errorf("expected spy sleeper to be called %d times got %d",3,sleeper.Calls)
	}
}