package handlers

import (
	"log"
	"net/http"

	"github.com/craniacshencil/got_to_do/internal/database"
	validation "github.com/craniacshencil/got_to_do/internal/signup"
	"github.com/craniacshencil/got_to_do/utils"
	"github.com/google/uuid"
)

const (
	usernameNotUniqueErr = "pq: duplicate key value violates unique constraint \"users_username_key\""
)

type SignupForm struct {
	Username        string
	FirstName       string
	LastName        string
	Password        string
	ConfirmPassword string
}

func (ApiConfig *ApiCfg) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupData SignupForm
	utils.ParseJSON(r, &signupData)
	if err := validation.ValidateSignup(signupData.Password, signupData.ConfirmPassword); err != nil {
		log.Println("ERR: Signup details are not valid.", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := ApiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Username:  signupData.Username,
		FirstName: signupData.FirstName,
		LastName:  signupData.LastName,
	})
	if err != nil {
		if err.Error() == usernameNotUniqueErr {
			log.Println(usernameNotUniqueErr)
			utils.WriteJSON(w, http.StatusInternalServerError, "Username is not unique")
			return
		}
		log.Println("Err: While creating user:", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}
