package handlers

import (
	"net/http"

	"github.com/craniacshencil/got_to_do/utils"
)

type LoginForm struct {
	Username string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData LoginForm
	utils.ParseJSON(r, &loginData)
	utils.WriteJSON(w, http.StatusCreated, "Success")
}
