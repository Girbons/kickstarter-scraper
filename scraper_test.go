package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestScraper(t *testing.T) {
	url := "https://www.kickstarter.com/projects/tabulagames/barbarians-the-invasion"
	result := ScrapeProject(url)

	assert.Equal(t, result.Creator, "Tabula Games")
	assert.Equal(t, result.AmountPledged, "£165,988")
	assert.Equal(t, result.AmountRequired, "£25,000")
	assert.Equal(t, result.Backers, "2,034 backers")
	assert.Equal(t, result.PledgeLevel[0], &PledgeLevel{"Barbarian Spirit", "£1"})
	assert.Equal(t, result.PledgeLevel[1], &PledgeLevel{"Wooden Edition", "£40"})
	assert.Equal(t, result.PledgeLevel[2], &PledgeLevel{"Iron Edition - EARLY BIRD", "£60"})
	assert.Equal(t, result.PledgeLevel[3], &PledgeLevel{"Iron Edition", "£67"})
	assert.Equal(t, result.PledgeLevel[4], &PledgeLevel{"Iron & Blood Edition (KS Exclusive)", "£85"})
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/kickstarter-project", ProjectScraper).Methods("GET")
	router.ServeHTTP(rr, req)

	return rr
}

func TestProjectScraperEndpoint(t *testing.T) {
	req, _ := http.NewRequest("GET", "/kickstarter-project?url=https://www.kickstarter.com/projects/tabulagames/barbarians-the-invasion", nil)
	response := executeRequest(req)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	result := `{"creator":"Tabula Games","amount_pledged":"£165,988","amount_required":"£25,000","backers":"2,034 backers","pledge_levels":[{"title":"Barbarian Spirit","amount":"£1"},{"title":"Wooden Edition","amount":"£40"},{"title":"Iron Edition - EARLY BIRD","amount":"£60"},{"title":"Iron Edition","amount":"£67"},{"title":"Iron \u0026 Blood Edition (KS Exclusive)","amount":"£85"}]}`

	assert.Equal(t, bodyString, result, "They must be equal")
}

func TestProjectScraperEndpointEmptyUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "/kickstarter-project?url=", nil)
	response := executeRequest(req)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	assert.Equal(t, bodyString, "accepted param is only url")
}

func TestProjectScraperEndpointNotKickstarter(t *testing.T) {
	req, _ := http.NewRequest("GET", "/kickstarter-project?url=http://google.com", nil)
	response := executeRequest(req)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	assert.Equal(t, bodyString, "Insert a kickstarter url")
}
