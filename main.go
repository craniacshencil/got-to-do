package main

import (
	"log"
	"net/http"
	"os"

	router "github.com/craniacshencil/got_to_do/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERR: couldn't load up env file", err)
	}

	clientRouter := router.MainRouter
	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: clientRouter,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ERROR:", err)
	}
}
