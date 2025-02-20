package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles(filepath.Join("templates", "template.html")))
}

func calculateEmissions(coalStr, fuelOilStr, naturalGasStr string) string {
	coalAmount, _ := strconv.ParseFloat(coalStr, 64)
	fuelOilAmount, _ := strconv.ParseFloat(fuelOilStr, 64)
	gasAmount, _ := strconv.ParseFloat(naturalGasStr, 64)

	emissionFactorCoal := 150.0
	emissionFactorFuelOil := 0.57
	emissionFactorGas := 0.0

	heatValueCoal := 20.47
	heatValueFuelOil := 40.40
	heatValueGas := 33.08

	totalCoalEmissions := (emissionFactorCoal * heatValueCoal * coalAmount) / 1_000_000
	totalFuelOilEmissions := (emissionFactorFuelOil * heatValueFuelOil * fuelOilAmount) / 1_000_000
	totalGasEmissions := (emissionFactorGas * gasAmount * heatValueGas) / 1_000_000

	totalEmissions := totalCoalEmissions + totalFuelOilEmissions + totalGasEmissions

	result := fmt.Sprintf(`Валові викиди при спалюванні палива:
Вугілля: %.4f т
Мазут: %.4f т
Природний газ: %.4f т
Загальна кількість викидів: %.4f т`,
		totalCoalEmissions, totalFuelOilEmissions, totalGasEmissions, totalEmissions)

	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		result = calculateEmissions(r.FormValue("coal"), r.FormValue("fuelOil"), r.FormValue("naturalGas"))
	}
	tmpl.Execute(w, result)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	fmt.Println("Сервер запущено на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
