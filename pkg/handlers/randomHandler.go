package handlers

import (
	"log"
	"net/http"
)

func (ApiConfig *ApiCfg) RandomHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		log.Println("Noo cookie:", err)
		return
	}
	log.Println("cookie-value: ", cookie)
	log.Println(cookie)
}
