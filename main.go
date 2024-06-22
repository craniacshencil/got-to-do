package main

import (
	"log"
	"net/http"
	"strconv"

	router "github.com/craniacshencil/got_to_do/router"
)

const (
	port = 8080
)

func main() {
	clientRouter := router.MainRouter
	server := http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: clientRouter,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ERROR:", err)
	}
}
