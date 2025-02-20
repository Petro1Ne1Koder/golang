package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles(filepath.Join("templates", "template.html")))
}

func calculateGaussian(x, mean, deviation float64) float64 {
	coefficient := 1 / (deviation * math.Sqrt(2*math.Pi))
	exponent := -math.Pow(x-mean, 2) / (2 * math.Pow(deviation, 2))
	return coefficient * math.Exp(exponent)
}

func approximateIntegral(start, end float64, intervals int, mean, deviation float64) float64 {
	sum := 0.0
	stepSize := (end - start) / float64(intervals)
	for i := 0; i < intervals; i++ {
		left := start + float64(i)*stepSize
		right := start + float64(i+1)*stepSize
		sum += (calculateGaussian(left, mean, deviation) + calculateGaussian(right, mean, deviation)) / 2 * stepSize
	}
	return sum
}

func performCalculation(power, initialError, improvedError, ratePerKWh float64) string {
	rangeStart := power - improvedError
	rangeEnd := power + improvedError
	divisions := 1000

	efficiencyBefore := approximateIntegral(rangeStart, rangeEnd, divisions, power, initialError)
	earningsBefore := power * 24 * efficiencyBefore * ratePerKWh * 1000
	penaltiesBefore := power * 24 * (1 - efficiencyBefore) * ratePerKWh * 1000

	efficiencyAfter := approximateIntegral(rangeStart, rangeEnd, divisions, power, improvedError)
	earningsAfter := power * 24 * efficiencyAfter * ratePerKWh * 1000
	penaltiesAfter := power * 24 * (1 - efficiencyAfter) * ratePerKWh * 1000

	result := fmt.Sprintf(`Прибуток до вдосконалення: %.2f тис. грн\nВиручка до вдосконалення: %.2f тис. грн\nШтраф до вдосконалення: %.2f тис. грн\nПрибуток після вдосконалення: %.2f тис. грн\nВиручка після вдосконалення: %.2f тис. грн\nШтраф після вдосконалення: %.2f тис. грн`,
		earningsBefore/1000,
		(earningsBefore-penaltiesBefore)/1000,
		penaltiesBefore/1000,
		earningsAfter/1000,
		(earningsAfter-penaltiesAfter)/1000,
		penaltiesAfter/1000)

	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err == nil {
			power, _ := strconv.ParseFloat(r.FormValue("inputPower"), 64)
			initialError, _ := strconv.ParseFloat(r.FormValue("firstErrorMargin"), 64)
			improvedError, _ := strconv.ParseFloat(r.FormValue("secondErrorMargin"), 64)
			ratePerKWh, _ := strconv.ParseFloat(r.FormValue("electricityRate"), 64)
			result = performCalculation(power, initialError, improvedError, ratePerKWh)
		}
	}
	tmpl.Execute(w, result)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	fmt.Println("Сервер запущено на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
