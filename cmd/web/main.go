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
	infoLogger := log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache()
	if err != nil {
		errLogger.Fatal(err)
	}

	startingUrls := make(map[string][]URL)

	app := application{
		urlsList:      startingUrls,
		templateCache: templateCache,
		infoLogger:    infoLogger,
		errLogger:     errLogger,
	}

	err = http.ListenAndServe(":8080", app.logRequest(secureHeaders(app.router())))

	if err != nil {
		errLogger.Fatal("Couldn't start server:", err)
	}

}
