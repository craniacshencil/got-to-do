package main

import (
	"log"
	"os"

	"github.com/craniacshencil/got_to_do/cmd/routes"
	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("APP_ENV")
	log.Println(env)
	if env == "docker" {
		if err := godotenv.Load(".env.docker"); err != nil {
			log.Println("Failed to load .env.docker: ", err)
		}
	} else {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Println("Failed to load .env.local: ", err)
		}
	}
}

func main() {
	routes.SetRoutes()
}
