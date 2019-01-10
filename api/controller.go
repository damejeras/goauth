package api

import (
	"fmt"
	"net/http"
)

func (s *server) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "goauth")
		if err != nil {
			panic(err)
		}
	}
}
