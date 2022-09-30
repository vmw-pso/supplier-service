package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/vmw-pso/supplier-service/data"
	"github.com/vmw-pso/toolkit"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
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
	fmt.Printf("Starting front-end, listening on :%d\n", *port)
	return http.ListenAndServe(addr, srv)
}

type server struct {
	DB     *sql.DB
	Models *data.Models
	mux    *chi.Mux
	tools  toolkit.Tools
}

func newServer() *server {
	mux := chi.NewMux()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	db := connectToDB()

	models := data.New()

	tools := toolkit.Tools{}

	srv := &server{
		DB:     db,
		mux:    mux,
		Models: models,
		tools:  tools,
	}
	srv.routes()

	return srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	counts := 0

	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Printf("Postgres not yet ready: %v\n", err)
			counts++
		} else {
			log.Println("Connected to PostgreSQL!")
			return conn
		}

		if counts > 10 {
			log.Println("Count not connect to database")
			return nil
		}

		log.Println("Backing off and waiting for 2 seconds...")
		time.Sleep(2 * time.Second)
	}
}
