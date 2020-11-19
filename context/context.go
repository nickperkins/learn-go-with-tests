package context

import (
	"context"
	"fmt"
	"net/http"
)

// A Store contains our data
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server will fetch data from our store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}
		fmt.Fprint(w, data)
	}
}
