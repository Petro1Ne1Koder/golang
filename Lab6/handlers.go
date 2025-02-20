package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func MainHandler(w http.ResponseWriter, r *http.Request) {
	data := performCalculations()
	templates.ExecuteTemplate(w, "main.html", data)
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	data := performCalculations()
	templates.ExecuteTemplate(w, "results.html", data)
}

func performCalculations() map[string]interface{} {
	type EP struct {
		Name       string
		Efficiency float64
		CosPhi     float64
		Voltage    float64
		Quantity   int
		Power      float64
		KV         float64
		TgPhi      float64
	}

	EPList := []EP{
		{"Шліфувальний верстат", 0.92, 0.9, 0.38, 4, 20, 0.15, 1.33},
		{"Свердлильний верстат", 0.92, 0.9, 0.38, 2, 14, 0.12, 1.00},
		{"Фугувальний верстат", 0.92, 0.9, 0.38, 4, 42, 0.15, 1.33},
		{"Циркулярна пила", 0.92, 0.9, 0.38, 1, 36, 0.3, 1.52},
		{"Прес", 0.92, 0.9, 0.38, 1, 20, 0.5, 0.75},
		{"Полірувальний верстат", 0.92, 0.9, 0.38, 1, 40, 0.2, 1.00},
		{"Фрезерний верстат", 0.92, 0.9, 0.38, 2, 32, 0.2, 1.00},
		{"Вентилятор", 0.92, 0.9, 0.38, 1, 20, 0.65, 0.75},
	}

	var totalPower, totalKVPower, totalKVPowerTg, totalPowerSquare float64
	for _, ep := range EPList {
		Pn := float64(ep.Quantity) * ep.Power
		totalPower += Pn
		totalKVPower += Pn * ep.KV
		totalKVPowerTg += Pn * ep.KV * ep.TgPhi
		totalPowerSquare += float64(ep.Quantity) * math.Pow(ep.Power, 2)
	}

	Kv := totalKVPower / totalPower
	ne := math.Round(math.Pow(totalPower, 2) / totalPowerSquare)
	Kr := 1.25
	Pp := Kr * totalKVPower
	Qp := Kv * totalPower * 1.33
	Sp := math.Sqrt(Pp*Pp + Qp*Qp)
	Ip := Pp / 0.38

	return map[string]interface{}{
		"Kv":     fmt.Sprintf("%.4f", Kv),
		"ne":     fmt.Sprintf("%.0f", ne),
		"Kr":     fmt.Sprintf("%.2f", Kr),
		"Pp":     fmt.Sprintf("%.2f", Pp),
		"Qp":     fmt.Sprintf("%.2f", Qp),
		"Sp":     fmt.Sprintf("%.2f", Sp),
		"Ip":     fmt.Sprintf("%.2f", Ip),
		"EPList": EPList,
	}
}
