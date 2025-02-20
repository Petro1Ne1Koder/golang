package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

var singlePhaseTemplate *template.Template

func init() {
	singlePhaseTemplate = template.Must(template.ParseFiles(filepath.Join("templates", "single_phase.html")))
}

func singlePhaseHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err == nil {
			voltage, _ := strconv.ParseFloat(r.FormValue("voltage"), 64)
			impedance, _ := strconv.ParseFloat(r.FormValue("impedance"), 64)
			if impedance != 0 {
				result = fmt.Sprintf("Струм однофазного КЗ: %.2f A", voltage/impedance)
			} else {
				result = "Помилка: Імпеданс не може бути нулем."
			}
		}
	}
	singlePhaseTemplate.Execute(w, result)
}
