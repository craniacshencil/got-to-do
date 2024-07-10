package handlers

import (
	"log"
	"net/http"

	"github.com/craniacshencil/got_to_do/pkg/myJwt"
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
	// MAKE CHANGES HERE TO SEND JWT TO FRONTEND
	tokenString, err := myJwt.CreateToken(username)
	if err != nil {
		log.Println("ERR: ", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	jwtCookie := &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(w, jwtCookie)

	utils.WriteJSON(w, http.StatusCreated, "Success")
}
