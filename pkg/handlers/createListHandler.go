package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/craniacshencil/got_to_do/pkg/myJwt"
	"github.com/craniacshencil/got_to_do/utils"
)

type Task struct {
	TaskName  string    `json:"task_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (ApiConfig *ApiCfg) CreateListHandler(w http.ResponseWriter, r *http.Request) {
	// Retreive jwt token from cookies
	token, err := r.Cookie("jwt")
	if err != nil {
		log.Println("ERR: Couldn't find cookie", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Validate the cookie
	_, err = myJwt.ValidateToken(token.Value)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Parse the json
	var Todo map[string]Task
	err = utils.ParseJSON(r, &Todo)
	if err != nil {
		log.Println("ERR: While parsing Todo JSON", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
