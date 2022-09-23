package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		port = flags.Int("port", 80, "port to listen on")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	addr := fmt.Sprintf(":%d", *port)

	srv := newServer()

	return http.ListenAndServe(addr, srv)
}

type server struct {
	mux *http.ServeMux
}

func newServer() *server {
	srv := &server{
		mux: http.NewServeMux(),
	}
	srv.routes()
	return srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) routes() {
	s.mux.Handle("/suppliers", s.handleSuppliers())
}

func (s *server) handleSuppliers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}
}
