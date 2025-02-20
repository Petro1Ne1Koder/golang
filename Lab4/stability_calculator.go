package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

var stabilityTemplate *template.Template

func init() {
	stabilityTemplate = template.Must(template.ParseFiles(filepath.Join("templates", "stability.html")))
}

func stabilityHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err == nil {
			current, _ := strconv.ParseFloat(r.FormValue("current"), 64)
			duration, _ := strconv.ParseFloat(r.FormValue("duration"), 64)
			result = fmt.Sprintf("Термічна стійкість: %.2f A²·с", current*current*duration)
		}
	}
	stabilityTemplate.Execute(w, result)
}
