package main

import (
	"net/http"
)

func (app *application) router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(200)

		// URL := "https://google.com"
		// statusCode := Ping(URL, 10)
		// toPrint := fmt.Sprintf("Status code for %s is %d", URL, statusCode)
		app.render(w, "home.html", pageData{URLs: app.urlsList})
	})

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	return router
}