package main

import (
	"log"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, temp string, data pageData) {
	ts, ok := app.templateCache[temp]
	if !ok {
		log.Printf("the template %s doesn't exist", temp)
		return
	}
	if err := ts.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.errLogger.Println(err)
	}
}
