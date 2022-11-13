package main

import (
	"log"
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

	router.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.WriteHeader(200)
			app.render(w, "newsite.html", pageData{URLs: app.urlsList})
		}
		if r.Method == http.MethodPost {
			form := r.ParseForm()
			if form != nil {
				log.Println("Couldn't parse form", form)
			}
			url := r.Form.Get("url")
			category := r.Form.Get("category")
			app.urlsList[category] = append(app.urlsList[category], URL{Url: url, StatusCode: 0})
			w.WriteHeader(200)
			app.render(w, "newsite.html", pageData{URLs: app.urlsList, Flash: "The website has been added successfully"})
		}
	})

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	return router
}
