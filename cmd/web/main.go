package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NBens/status/internal/helpers"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(200)
		URL := "https://google.com"
		statusCode := helpers.Ping(URL, 10)
		toPrint := fmt.Sprintf("Status code for %s is %d", URL, statusCode)
		w.Write([]byte(toPrint))
	})

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal("Couldn't start server:", err)
	}

}
