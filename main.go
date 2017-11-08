package main

import (
	"log"
	"net/http"

	"github.com/Girbons/kickstarter-scraper/scraper"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/kickstarter-project", scraper.ProjectScraper).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
