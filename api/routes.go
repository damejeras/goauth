package api

import (
	"github.com/damejeras/goauth/scopes"
	"github.com/damejeras/goauth/users"
)

func (s *server) routes() {
	s.router.Use(contentTypeJSON)
	s.router.HandleFunc("/register", users.RegistrationHandler()).Methods("POST")
	s.router.HandleFunc("/login", users.LoginHandler()).Methods("POST")

	scopesAPI := s.router.PathPrefix("/scopes").Subrouter()
	scopesAPI.HandleFunc("/{email}", scopes.HandleNew()).Methods("POST")
}
