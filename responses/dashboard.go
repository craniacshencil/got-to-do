package responses

import (
	"log"
	"net/http"
	"text/template"
)

const (
	dashboardHtml = "static/dash.html"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(dashboardHtml)
	if err != nil {
		log.Fatal("ERROR:", err)
	} else {
		t.Execute(w, nil)
	}
}
