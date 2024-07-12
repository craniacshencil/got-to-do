package handlers

import (
	"log"
	"net/http"

	"github.com/craniacshencil/got_to_do/pkg/myJwt"
	"github.com/craniacshencil/got_to_do/utils"
)

func (ApiConfig *ApiCfg) CreateListHandler(w http.ResponseWriter, r *http.Request) {
	// Retreive jwt token from cookies
	token, err := r.Cookie("jwt")
	if err != nil {
		log.Println("ERR: Couldn't find cookie", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO:Validate the cookie
	_, err = myJwt.ValidateToken(token.Value)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("Validation successful")

	// myJwt.ValidateToken(r.Cookie("jwt").Value)
}
