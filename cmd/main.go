package main

import (
	"fmt"
	data "groupie_upgrade/data"
	error "groupie_upgrade/error"
	handler "groupie_upgrade/handler"
	"net/http"
)

var port = ":5050"

func main() {
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	data.FetchData("https://groupietrackers.herokuapp.com/api/artists")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Reception de l'URL, 4 pages: home - artist - recherche - filtre
		if r.URL.Path == "/" || r.URL.Path == "" {
			handler.Home(w, r)
		} else if r.URL.Path == "/artist" {
			handler.Artist_handler(w, r)
		} else if r.URL.Path == "/searched_artist" {
			query := r.FormValue("search_bar")
			handler.Searched_artist(w, r, query)
		} else if r.URL.Path == "/filtered_artists" {
			handler.Filtered_artists(w, r)
		} else {
			error.Error_handler(w, r, 404)
		}
	})
	// link URL
	fmt.Println("http://localhost:5050/")
	http.ListenAndServe(port, nil)
}
