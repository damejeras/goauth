package api

func (s *server) routes() {
	s.router.HandleFunc("/", s.get()).Methods("GET")
}
