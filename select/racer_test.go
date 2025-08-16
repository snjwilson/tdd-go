package learnselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test the speed of the servers, returning the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(10*time.Millisecond)
		fastServer := makeDelayedServer(0*time.Millisecond)
	
		defer slowServer.Close()
		defer fastServer.Close()
	
		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got, err := Racer(slowURL, fastURL)
		
		if err != nil {
			t.Errorf("expected no error but got error %v", err.Error())
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("test error if there is a delay more than timeout", func(t *testing.T) {
		delay :=20*time.Millisecond
		slowServer := makeDelayedServer(delay)
		slowerServer := makeDelayedServer(delay)

		defer slowServer.Close()
		defer slowerServer.Close()

		slowServerURL := slowServer.URL
		slowerServerURL := slowerServer.URL

		_, err := ConfigurableRacer(slowServerURL, slowerServerURL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but did not receive any")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}