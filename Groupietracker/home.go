package Groupie

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./templates/home.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
}
