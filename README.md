# Kickstarter Scraper

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
