package handlers

import (
	"log"
	"net/http"

	"github.com/craniacshencil/got_to_do/pkg/passwords"
	"github.com/craniacshencil/got_to_do/utils"
)

type LoginForm struct {
	Username string
	Password string
}

func (ApiConfig *ApiCfg) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData LoginForm
	utils.ParseJSON(r, &loginData)
	username := loginData.Username
	password := loginData.Password
	usernameAndPassword, err := ApiConfig.DB.GetUsernameAndPassword(r.Context(), username)
	if err != nil {
		log.Println("ERR: While retreiving Username and Password", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = passwords.MatchPassword(password, usernameAndPassword.Password); err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(usernameAndPassword)

	utils.WriteJSON(w, http.StatusCreated, "Success")
}
