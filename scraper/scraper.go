package scraper

import (
	"encoding/json"
	"fmt"

	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"

	"net/http"
	"strings"
)

// Project info to retrieve.
type Project struct {
	Creator        string         `json:"creator"`
	AmountPledged  string         `json:"amount_pledged"`
	AmountRequired string         `json:"amount_required"`
	Backers        string         `json:"backers"`
	PledgeLevel    []*PledgeLevel `json:"pledge_levels"`
}

// PledgeLevel info to retrieve.
type PledgeLevel struct {
	Title  string `json:"title"`
	Amount string `json:"amount"`
}

// ParseLevel Retrieve the core info of a PledgeLevel.
func ParseLevel(node *html.Node) *PledgeLevel {
	amount := htmlquery.InnerText(htmlquery.FindOne(node, "//h2[@class='pledge__amount']//span[@class='money']/text()"))
	title := strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(node, "//h3[@class='pledge__title']/text()")))
	return &PledgeLevel{title, amount}
}

// ScrapeProject parse request.Body and extract the required info.
func ScrapeProject(url string) *Project {
	doc, err := htmlquery.LoadURL(url)

	if err != nil {
		panic(err)
	}

	project := Project{}

	project.Creator = strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='creator-name']//div//a/text()")))
	project.AmountPledged = htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='NS_campaigns__spotlight_stats']//span/text()"))
	project.Backers = htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='NS_campaigns__spotlight_stats']//b/text()"))
	project.AmountRequired = htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='type-12 medium navy-500']//span[@class='money']//span/text()"))

	for _, level := range htmlquery.Find(doc, "//div[@class='pledge__info']") {
		project.PledgeLevel = append(project.PledgeLevel, ParseLevel(level))
	}
	return &project
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
