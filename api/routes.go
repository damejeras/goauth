package api

import "github.com/damejeras/goauth/users"

func (s *server) routes() {
	s.router.Use(jsonAPI)

	// TODO: rename this
	userApi := s.router.PathPrefix("/users").Subrouter()
	userApi.HandleFunc("/register", users.RegistrationHandler()).Methods("POST")
	userApi.HandleFunc("/login", users.LoginHandler()).Methods("POST")
}
