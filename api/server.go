package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

// NewServer creates new api.server instance
func NewServer() *server {
	s := server{mux.NewRouter()}
	s.routes()
	return &s
}

// ServeHTTP is used to server requests to default HTTP server
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
