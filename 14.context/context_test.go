package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		// Got cancelled, so return an error and empty result
		log.Println("got cancelled so tossing", data)
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

// Implementing the http.ResponseWriter interface
func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

// Implementing the http.ResponseWriter interface
func (s *SpyResponseWriter) Write(buf []byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

// Implementing the http.ResponseWriter interface
func (s *SpyResponseWriter) WriteHeader(code int) {
	s.written = true
}

/*func (s *SpyStore) assertCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}
*/

func TestServer(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if got := response.Body.String(); got != data {
			t.Errorf("got %q, want %q", got, data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// Get a new Done channel for the requests context
		cancellingCtx, cancel := context.WithCancel(request.Context())
		// Send an async call to cancel after 5 ms
		time.AfterFunc(5*time.Millisecond, cancel)
		// Update the context with the cancelling
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}

	})

	/*
		t.Run("cancel work if request is cancelled", func(t *testing.T) {
			data := "hello, world"
			store := &SpyStore{data, false, t}
			svr := Server(store)

			request := httptest.NewRequest(http.MethodGet, "/", nil)

			cancellingCtx, cancel := context.WithCancel(request.Context())
			time.AfterFunc(5*time.Millisecond, cancel)
			request = request.WithContext(cancellingCtx)

			response := httptest.NewRecorder()

			svr.ServeHTTP(response, request)

			store.assertCancelled()
		})*/
}
