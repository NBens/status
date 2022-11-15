package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type URL struct {
	Url        string
	StatusCode int
}

type application struct {
	urlsList      map[string][]URL
	templateCache map[string]*template.Template
	infoLogger    *log.Logger
	errLogger     *log.Logger
}

func main() {

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	startingUrls := make(map[string][]URL)

	infoLogger := log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
	app := application{
		urlsList:      startingUrls,
		templateCache: templateCache,
		infoLogger:    infoLogger,
		errLogger:     errLogger,
	}

	err = http.ListenAndServe(":8080", app.router())

	if err != nil {
		log.Fatal("Couldn't start server:", err)
	}

}
