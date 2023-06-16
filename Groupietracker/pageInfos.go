package Groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func Infos(w http.ResponseWriter, r *http.Request) {
	tmpl2 := template.Must(template.ParseFiles("./templates/infos.html"))
	var oneartist Artist
	var l OneLocations
	id := r.FormValue("id")
	request, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	request2, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	request3, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(request.Body)
	byteValue2, _ := ioutil.ReadAll(request2.Body)
	byteValue3, _ := ioutil.ReadAll(request3.Body)
	request.Body.Close()
	request2.Body.Close()
	request3.Body.Close()
	json.Unmarshal(byteValue, &oneartist)
	json.Unmarshal(byteValue2, &oneartist.Relation)
	json.Unmarshal(byteValue3, &l)
	oneartist.Locations = l.Locations
	tmpl2.Execute(w, oneartist)
	fmt.Println(oneartist.Locations)
}
