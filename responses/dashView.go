package responses

import (
	"html/template"
	"log"
	"net/http"
)

const (
	dashboardHtml = "static/dash.html"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	dashTemplate, err := template.ParseFiles(dashboardHtml)
	if err != nil {
		log.Fatal("ERROR:", err)
	} else {
		dashTemplate.Execute(w, nil)
	}
}
