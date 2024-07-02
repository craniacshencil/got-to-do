package routes

import (
	"log"
	"net/http"

	pkg "github.com/craniacshencil/got_to_do/pkg/signup"
	"github.com/craniacshencil/got_to_do/utils"
	"github.com/go-chi/chi/v5"
)

func SetRoutes() {
	MainRouter := chi.NewRouter()
	MainRouter.Get("/", simplePing)
	MainRouter.Post("/", simplePing)
	MainRouter.Put("/", simplePing)
	MainRouter.Post("/signup", pkg.SignupHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: MainRouter,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println("ERR: Starting the server")
	}
}

func simplePing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		utils.WriteJSON(w, 200, "POST-y business")
	} else if r.Method == "GET" {
		utils.WriteJSON(w, 200, "Get request essketit")
	} else {
		utils.WriteJSON(w, 200, "Weird Method Ew")
	}
}
