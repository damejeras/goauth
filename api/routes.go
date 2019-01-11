package api

import "github.com/damejeras/goauth/users"

func (s *server) routes() {
	s.router.Use(jsonAPI)

	s.router.HandleFunc("/", s.get()).Methods("GET")

	// TODO: rename this
	userApi := s.router.PathPrefix("/users").Subrouter()
	userApi.HandleFunc("/", users.RegistrationHandler()).Methods("POST")
	userApi.HandleFunc("/login", users.LoginHandler()).Methods("POST")
}
