package main

import (
	"html/template"
	"log"
	"net/http"
)

type application struct {
	urlsList      []string
	templateCache map[string]*template.Template
}

func main() {

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app := application{
		urlsList:      []string{"https://koora.com", "https://google.com"},
		templateCache: templateCache,
	}

	err = http.ListenAndServe(":8080", app.router())

	if err != nil {
		log.Fatal("Couldn't start server:", err)
	}

}
