package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/NBens/status/internal/helpers"
)

type pageData struct {
	Title   string
	Content string
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(200)
		ts, err := template.ParseFiles("./ui/html/index.html")
		if err != nil {
			log.Fatal("Couldn't parse templates", err)
		}

		URL := "https://google.com"
		statusCode := helpers.Ping(URL, 10)
		toPrint := fmt.Sprintf("Status code for %s is %d", URL, statusCode)

		ts.Execute(w, pageData{Title: "Home Page", Content: toPrint})
	})

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal("Couldn't start server:", err)
	}

}
