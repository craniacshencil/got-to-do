package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const (
    port = 8080
)

func dashboard(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("index.html")
    if err != nil {
        log.Fatal("ERROR:", err)
    } else {
        t.Execute(w, nil)
    }
    
}

func main() {
    r := chi.NewRouter()

    server := http.Server{
        Addr: ":" + strconv.Itoa(port),
        Handler: r,
    }

    r.Get("/", dashboard)

    if err := server.ListenAndServe(); err != nil {
        log.Fatal("ERROR:", err)
    }
}
