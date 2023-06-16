package Groupie

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Filters(filters *Filter, w http.ResponseWriter, r *http.Request) {
	filters.resfilter = []Artist{}
	FilterRadio(filters, r.FormValue("memberfilter"))
	FilterLocation(filters, r.FormValue("country"))
	Rangebyfirstalbum(filters, r.FormValue("firstalbum"))
	Rangebycreationdate(filters, r.FormValue("creationdate"))
}

func FilterRadio(filters *Filter, form string) {
	for i := 0; i < len(filters.filterartist); i++ {
		if form == strconv.Itoa(len(filters.filterartist[i].Members)) {
			fmt.Print(filters.filterartist[i].Name)
			filters.resfilter = append(filters.resfilter, filters.filterartist[i])
		}
	}
}

func FilterLocation(filters *Filter, form string) {
	var countryfilter []string
	for i := 0; i < len(filters.filterartist); i++ {
		for j := 0; j < len(filters.filterartist[i].Locations); j++ {
			countryfilter = append(countryfilter, strings.Split(filters.filterartist[i].Locations[j], "-")[1])
		}
		for k := 0; k < len(countryfilter); k++ {
			if countryfilter[k] == form {
				fmt.Print(filters.filterartist[i].Name)
				filters.resfilter = append(filters.resfilter, filters.filterartist[i])
			}
		}
	}
}

func Rangebyfirstalbum(filters *Filter, form string) {
	for i := 0; i < len(filters.filterartist); i++ {
		if form == strings.Split(filters.filterartist[i].FirstAlbum, "-")[2] {
			fmt.Print(filters.filterartist[i].Name)
			filters.resfilter = append(filters.resfilter, filters.filterartist[i])
		}
	}
}

func Rangebycreationdate(filters *Filter, form string) {

	for i := 0; i < len(filters.filterartist); i++ {
		if form == strconv.Itoa(filters.filterartist[i].CreationDate) {
			fmt.Print(filters.filterartist[i].Name)
			filters.resfilter = append(filters.resfilter, filters.filterartist[i])
		}
	}
}
