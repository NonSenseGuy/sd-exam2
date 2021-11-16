package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nonsenseguy/sd-exam2/handlers"
	"github.com/nonsenseguy/sd-exam2/store"
)

func main() {
	args := Args{
		conn: "postgres://postgres:@localhost:5432/posgres?sslmode=disable",
		port: "8080",
	}

	if conn := os.Getenv("DB_CONN"); conn != "" {
		args.conn = conn
	}

	if port := os.Getenv("PORT"); port != "" {
		args.port = ":" + port
	}

	if err := Run(args); err != nil {
		log.Println(err)
	}
}

type Args struct {
	conn string
	port string
}

func Run(args Args) error {
	router := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()

	st := store.NewPostgresConnection(args.conn)
	handler := handlers.NewFraudataHandler(st)
	RegisterAllRoutes(router, handler)

	log.Println("Starting server at port: ", args.port)

	return http.ListenAndServe(args.port, router)
}

func RegisterAllRoutes(router *mux.Router, handler handlers.IFraudataHandler) {
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/fraudata/item", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/fraudata", handler.Report).Methods(http.MethodPost)
	router.HandleFunc("/fraudata", handler.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/fraudata", handler.List).Methods(http.MethodGet)
	router.HandleFunc("/health", handler.Health).Methods(http.MethodGet)
}
