package responses

import (
	"html/template"
	"log"
	"net/http"

	validation "github.com/craniacshencil/got_to_do/validation"
)

const (
	signupHtml = "static/signup.html"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		signupTemplate, err := template.ParseFiles(signupHtml)
		if err != nil {
			log.Fatal("ERROR:", err)
		} else {
			signupTemplate.Execute(w, nil)
		}
	} else {
		err := validation.ValidateSignup(r)
		if err != nil {
			log.Println("ERROR:", err)
		} else {
			// Database Interaction
		}
	}
}
