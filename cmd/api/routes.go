package main

func (s *server) routes() {
	s.mux.Post("/suppliers", s.handleSupplierNames())
}
