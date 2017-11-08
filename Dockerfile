FROM golang:latest

MAINTAINER Alessandro De Angelis <alessandrodea22@gmail.com>

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get github.com/gorilla/mux
RUN go get github.com/antchfx/xquery/html
RUN go get github.com/girbons/kickstarter-scraper

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
