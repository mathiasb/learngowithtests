package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data, err := store.Fetch(req.Context())

		if err != nil {
			// Just terminate with some message and leave the response empty
			return
		}

		fmt.Fprint(w, data)
	}
}
