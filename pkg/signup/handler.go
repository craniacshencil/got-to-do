package pkg

import (
	"log"
	"net/http"

	validation "github.com/craniacshencil/got_to_do/internal/signup"
	"github.com/craniacshencil/got_to_do/utils"
)

type SignupForm struct {
	Username        string
	FirstName       string
	LastName        string
	Password        string
	ConfirmPassword string
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupData SignupForm
	utils.ParseJSON(r, &signupData)
	if err := validation.ValidateSignup(signupData.Password, signupData.ConfirmPassword); err != nil {
		log.Println("ERR: Signup details are not valid.", err)
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusCreated, "Success")
}
