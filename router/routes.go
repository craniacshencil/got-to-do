package router

import (
	requests "github.com/craniacshencil/got_to_do/requests"
	responses "github.com/craniacshencil/got_to_do/responses"
	"github.com/go-chi/chi/v5"
)

var MainRouter *chi.Mux

func init() {
	MainRouter = chi.NewRouter()
	MainRouter.Get("/", responses.Dashboard)
	MainRouter.Get("/signup", responses.SignupGET)
	MainRouter.Post("/signup", responses.SignupPOST)

	usersRouter := chi.NewRouter()
	usersRouter.Post("/createAccount", requests.CreateAccount)

	MainRouter.Mount("/users", usersRouter)
}
