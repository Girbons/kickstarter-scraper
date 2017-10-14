package main

import (
	"fmt"
	"testing"
)

func TestScraper(t *testing.T) {
	url := "https://www.kickstarter.com/projects/tabulagames/barbarians-the-invasion"
	result := ScrapeProject(url)
	fmt.Printf("%+v", result)
	if (result != Project{"Tabula Games", "£165,988", "2,034 backers"}) {
		t.Errorf("Result was incorrect, %v", result)
	}
}
