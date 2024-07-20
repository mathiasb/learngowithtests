package selecter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer slowServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("did not expect an error but got %s", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("timeout when longer that a threshold", func(t *testing.T) {
		slowServer := makeDelayedServer(2 * time.Second)
		fastServer := makeDelayedServer(3 * time.Second)

		defer slowServer.Close()
		defer slowServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 1*time.Second)

		if err == nil {
			t.Error("Expected an error but got nil")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
