package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
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
	mu            *sync.Mutex
}

func main() {
	infoLogger := log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache()
	if err != nil {
		errLogger.Fatal(err)
	}

	startingUrls := make(map[string][]URL)
	startingUrls["Demo"] = append(startingUrls["Demo"], URL{Url: "https://google.com", StatusCode: 0}) // Adding demo data for testing

	app := application{
		urlsList:      startingUrls,
		templateCache: templateCache,
		infoLogger:    infoLogger,
		errLogger:     errLogger,
	}

	port := flag.Int("port", 8080, "The application port")

	flag.Parse()

	err = http.ListenAndServe(":"+strconv.Itoa(*port), app.logRequest(secureHeaders(app.router())))

	if err != nil {
		errLogger.Fatal("Couldn't start server:", err)
	}

	for {
		for _, value := range app.urlsList {
			for _, url := range value {
				go func(rl *URL) {
					statusCode := Ping(rl.Url, 5)
					app.mu.Lock()
					rl.StatusCode = statusCode
					app.mu.Unlock()
					app.infoLogger.Println("Scraping", rl.Url)
				}(&url)
			}
		}
		time.Sleep(60 * time.Second)
	}

}
