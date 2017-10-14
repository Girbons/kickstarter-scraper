# Kickstarter Scraper

[![Build Status](https://travis-ci.org/Girbons/kickstarter-scraper.svg?branch=master)](https://travis-ci.org/Girbons/kickstarter-scraper)
[![Go Report Card](https://goreportcard.com/badge/github.com/girbons/kickstarter-scraper)](https://goreportcard.com/report/github.com/girbons/kickstarter-scraper)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Install

```
go get github.com/gorilla/mux
go get github.com/antchfx/xquery/html
```

## Run project

`go run main.go`


### Example request

```
http://localhost:8080/kickstarter-project?url=[project url]

return JSON response
```


### Docker

```
docker build -t kickstarter-scraper .
docker run -p 8080:8080 kickstarter-scraper
```
