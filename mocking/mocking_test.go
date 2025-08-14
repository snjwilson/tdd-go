package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpyCountdownOps struct {
	Calls	[]string
}

func (s *SpyCountdownOps) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOps) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpyCountdownOps{}
		Countdown(buffer, spy)
		got := buffer.String()
		expected := `3
2
1
Go!`
		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	})
	
	t.Run("calls sleep before print", func(t *testing.T) {
		spy := &SpyCountdownOps{}
		Countdown(spy, spy)
		want := []string{write,sleep,write,sleep,write,sleep,write}
		if !reflect.DeepEqual(spy.Calls, want) {
			t.Errorf("expected %v got %v", want, spy.Calls)
		}
	})
}