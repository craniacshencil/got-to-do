package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	validation "github.com/craniacshencil/got_to_do/validation"
	"github.com/joho/godotenv"
)

const (
	signupHtml = "static/signup.html"
)

type SignupForm struct {
	Username        string `json:"username"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func SignupGET(w http.ResponseWriter, r *http.Request) {
	signupTemplate, err := template.ParseFiles(signupHtml)
	if err != nil {
		log.Fatal("ERROR:", err)
	} else {
		signupTemplate.Execute(w, nil)
	}
}

func SignupPOST(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERR: couldn't load up env file", err)
	}

	err = validation.ValidateSignup(r)
	if err != nil {
		log.Println("ERROR:", err)
	} else {
		signupDetails := &SignupForm{
			Username:        r.FormValue("username"),
			FirstName:       r.FormValue("firstName"),
			LastName:        r.FormValue("lastName"),
			Password:        r.FormValue("password"),
			ConfirmPassword: r.FormValue("confirmPassword"),
		}

		signupJSON, err := json.Marshal(signupDetails)
		if err != nil {
			log.Fatal("Messed up struct")
		}
		signupReadable := bytes.NewReader(signupJSON)

		postURL := fmt.Sprintf("http://localhost:%s/users/createAccount", os.Getenv("PORT"))
		http.Post(postURL, "application/json", signupReadable)
	}
}
