package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
)

var threePhaseTemplate *template.Template

func init() {
	threePhaseTemplate = template.Must(template.ParseFiles(filepath.Join("templates", "three_phase.html")))
}

func threePhaseHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err == nil {
			voltage, _ := strconv.ParseFloat(r.FormValue("voltage"), 64)
			impedance, _ := strconv.ParseFloat(r.FormValue("impedance"), 64)
			if impedance != 0 {
				result = fmt.Sprintf("Струм трифазного КЗ: %.2f A", voltage/(impedance*math.Sqrt(3)))
			} else {
				result = "Помилка: Імпеданс не може бути нулем."
			}
		}
	}
	threePhaseTemplate.Execute(w, result)
}
