package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Главная страница
func MainHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "main.html", nil)
}

// Однофазный КЗ
func SinglePhaseHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		voltage, _ := strconv.ParseFloat(r.FormValue("voltage"), 64)
		impedance, _ := strconv.ParseFloat(r.FormValue("impedance"), 64)
		if impedance != 0 {
			result = fmt.Sprintf("Струм однофазного КЗ: %.2f A", voltage/impedance)
		} else {
			result = "Помилка: Імпеданс не може бути нулем."
		}
	}
	templates.ExecuteTemplate(w, "single_phase.html", result)
}

// Трифазный КЗ
func ThreePhaseHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		voltage, _ := strconv.ParseFloat(r.FormValue("voltage"), 64)
		impedance, _ := strconv.ParseFloat(r.FormValue("impedance"), 64)
		if impedance != 0 {
			result = fmt.Sprintf("Струм трифазного КЗ: %.2f A", voltage/(impedance*math.Sqrt(3)))
		} else {
			result = "Помилка: Імпеданс не може бути нулем."
		}
	}
	templates.ExecuteTemplate(w, "three_phase.html", result)
}

// Проверка устойчивости
func StabilityHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		current, _ := strconv.ParseFloat(r.FormValue("current"), 64)
		duration, _ := strconv.ParseFloat(r.FormValue("duration"), 64)
		result = fmt.Sprintf("Термічна стійкість: %.2f A²·с", current*current*duration)
	}
	templates.ExecuteTemplate(w, "stability.html", result)
}
