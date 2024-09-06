package handlers

import (
	"log"
	"net/http"

	"github.com/craniacshencil/got_to_do/internal/database"
	validation "github.com/craniacshencil/got_to_do/internal/signup"
	"github.com/craniacshencil/got_to_do/pkg/passwords"
	"github.com/craniacshencil/got_to_do/utils"
	"github.com/google/uuid"
)

const (
	usernameNotUniqueErr = "pq: duplicate key value violates unique constraint \"users_username_key\""
)

type SignupForm struct {
	Username        string `json:"username"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (ApiConfig *ApiCfg) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupData SignupForm
	utils.ParseJSON(r, &signupData)
	if err := validation.ValidateSignup(signupData.Password, signupData.ConfirmPassword); err != nil {
		log.Println("ERR: Signup details are not valid.", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Send user details to database

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

	// Send password to database

	hash, err := passwords.HashPassword(signupData.Password)
	if err != nil {
		log.Println("ERR:", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	password, err := ApiConfig.DB.CreatePassword(r.Context(), database.CreatePasswordParams{
		ID:       user.ID,
		Password: string(hash),
	})
	if err != nil {
		log.Println("ERR: While storing password in DB:", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusCreated, password)
}
