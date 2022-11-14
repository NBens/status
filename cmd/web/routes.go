package main

import (
	"net/http"
)

func (app *application) router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", app.home)

	router.HandleFunc("/add", app.add)

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	return router
}
