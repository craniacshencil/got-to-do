package responses

import (
	"log"
	"net/http"
	"text/template"
)

const (
	dashboardHtml = "static/dash.html"
	signupHtml = "static/signup.html"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	dashTemplate, err := template.ParseFiles(dashboardHtml)
	if err != nil {
		log.Fatal("ERROR:", err)
	} else {
		dashTemplate.Execute(w, nil)
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	signupTemplate, err := template.ParseFiles(signupHtml)
	if err != nil {
		log.Fatal("ERROR:", err)
	} else {
		signupTemplate.Execute(w, nil)
	}
}
