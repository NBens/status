package main

import (
	"html/template"
	"log"
	"path/filepath"
)

type pageData struct {
	URLs  map[string][]URL
	Flash string
}

func IsLast(i, size int) bool {
	return i == size-1
}

var functions = template.FuncMap{
	"IsLast": IsLast,
}

func newTemplateCache() (map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)
	files, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := filepath.Base(file)
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(file)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
