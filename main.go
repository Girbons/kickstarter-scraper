package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log"
	"strings"

	"github.com/antchfx/xquery/html"
	"github.com/gorilla/mux"
)

// Project info to retrieve.
type Project struct {
	Creator string `json:"creator"`
	Amount  string `json:"amount"`
	Backers string `json:"backers"`
}

// ScrapeProject parse request.Body and extract the required info.
func ScrapeProject(url string) Project {
	doc, err := htmlquery.LoadURL(url)

	if err != nil {
		panic(err)
	}

	creator := htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='creator-name']//div//a/text()"))
	amount := htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='NS_campaigns__spotlight_stats']//span/text()"))
	backers := htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='NS_campaigns__spotlight_stats']//b/text()"))

	return Project{strings.TrimSpace(creator), amount, backers}

}

// ProjectScraper return the json response.
func ProjectScraper(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	url := vars.Get("url")

	if url == "" {
		fmt.Fprintf(w, "accepted param is only url")
		return
	}

	if strings.HasPrefix(url, "https://www.kickstarter.com") {
		data := ScrapeProject(url)
		outgoingJSON, _ := json.Marshal(data)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(outgoingJSON))
	} else {
		fmt.Fprintf(w, "Insert a kickstarter url")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/kickstarter-project", ProjectScraper).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
