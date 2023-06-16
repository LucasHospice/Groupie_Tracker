package Groupie

import (
	"net/http"
)

type Filter struct {
	filterartist   []Artist
	resfilter      []Artist
	filterlocation []OneLocations
}

type OneRelation struct {
	Id             int
	DatesLocations map[string][]string
}

type OneLocations struct {
	Id        int
	Locations []string
}

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	Relation     OneRelation
}

func Server() {
	filters := &Filter{}
	http.HandleFunc("/", Home)
	http.HandleFunc("/artistpage/infos", Infos)
	http.HandleFunc("/artistpage", func(w http.ResponseWriter, r *http.Request) {
		Artistpage(filters, w, r)
	})
	http.HandleFunc("/filter", func(w http.ResponseWriter, r *http.Request) {
		Filters(filters, w, r)
		http.Redirect(w, r, "/artistpage", http.StatusFound)
	})
	http.HandleFunc("/contact", Contact)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
