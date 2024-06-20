package main

import (
	"log"
	"net/http"
	"strconv"

	responses "github.com/craniacshencil/got_to_do/responses"

	"github.com/go-chi/chi/v5"
)

const (
	port = 8080
)

func main() {
	r := chi.NewRouter()

	server := http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: r,
	}

	r.Get("/", responses.Dashboard)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ERROR:", err)
	}
}
