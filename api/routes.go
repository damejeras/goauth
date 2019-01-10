package api

func (s *server) routes() {
	s.router.Use(jsonAPI)

	s.router.HandleFunc("/", s.get()).Methods("GET")
}
