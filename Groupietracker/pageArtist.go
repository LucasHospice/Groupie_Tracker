package Groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func Artistpage(filters *Filter, w http.ResponseWriter, r *http.Request) {
	var artists []Artist
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	request, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()
	json.Unmarshal(body, &artists)
	filters.filterartist = artists
	//CHECK FILTER
	// APPLY FILTER -> SUPPRIMER LES ELEMENTS QUI CORRESPONDENT PAS
	if len(filters.resfilter) > 0 {
		artists = filters.resfilter
	}

	tmpl.Execute(w, artists)
}
