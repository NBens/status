package main

import (
	"html/template"
	"log"
	"net/http"
)

type URL struct {
	Url        string
	StatusCode int
}

type application struct {
	urlsList      map[string][]URL
	templateCache map[string]*template.Template
}

func main() {

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	startingUrls := make(map[string][]URL)
	startingUrls["Category Something"] = []URL{URL{Url: "https://koora.com", StatusCode: 200}}
	app := application{
		urlsList:      startingUrls,
		templateCache: templateCache,
	}

	err = http.ListenAndServe(":8080", app.router())

	if err != nil {
		log.Fatal("Couldn't start server:", err)
	}

}
