package routes

import (
	"log"
	"net/http"

	"github.com/craniacshencil/got_to_do/pkg/handlers"
	"github.com/craniacshencil/got_to_do/utils"
	"github.com/go-chi/chi/v5"
)

func SetRoutes() {
	MainRouter := chi.NewRouter()
	apiConfig := handlers.SetupDB()
	/* MainRouter.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"Set-Cookie",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           3600, // Maximum value not ignored by any of major browsers
	})) */
	MainRouter.Get("/", simplePing)
	MainRouter.Post("/", simplePing)
	MainRouter.Put("/", simplePing)
	MainRouter.Post("/signup", apiConfig.SignupHandler)
	MainRouter.Post("/login", apiConfig.LoginHandler)

	MainRouter.Get("/random", apiConfig.RandomHandler)

	UserRouter := chi.NewRouter()
	UserRouter.Post("/", apiConfig.CreateListHandler)
	UserRouter.Get("/", apiConfig.DisplayListHandler)
	UserRouter.Post("/{list_id}", apiConfig.AddToListHandler)
	UserRouter.Put("/{list_id}", apiConfig.UpdateListHandler)
	UserRouter.Delete("/{list_id}", apiConfig.DeleteListHandler)
	UserRouter.Delete("/task/{task_id}", apiConfig.DeleteTaskHandler)
	MainRouter.Mount("/users/{user_id}/{date}", UserRouter)

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
