package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestScraper(t *testing.T) {
	url := "https://www.kickstarter.com/projects/tabulagames/barbarians-the-invasion"
	result := ScrapeProject(url)
	if (result != Project{"Tabula Games", "£165,988", "2,034 backers"}) {
		t.Errorf("Result was incorrect, %v", result)
	}
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
	result := strings.TrimSpace(`{"creator":"Tabula Games","amount":"£165,988","backers":"2,034 backers"}`)
	if bodyString != result {
		t.Errorf("Incorrect response, %v", response.Body)
	}
}

func TestProjectScraperEndpointEmptyUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "/kickstarter-project?url=", nil)
	response := executeRequest(req)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	result := "accepted param is only url"
	if bodyString != result {
		t.Errorf("Incorrect response, %v", response.Body)
	}
}

func TestProjectScraperEndpointNotKickstarter(t *testing.T) {
	req, _ := http.NewRequest("GET", "/kickstarter-project?url=http://google.com", nil)
	response := executeRequest(req)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	result := "Insert a kickstarter url"
	if bodyString != result {
		t.Errorf("Incorrect response, %v", response.Body)
	}
}
