package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func MainHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "main.html", nil)
}

func Task1Handler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		elements := r.FormValue("elements")
		n, _ := strconv.ParseFloat(r.FormValue("nValue"), 64)
		result = calculateTask1(elements, n)
	}
	templates.ExecuteTemplate(w, "task1.html", result)
}

func Task2Handler(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.Method == http.MethodPost {
		omega, _ := strconv.ParseFloat(r.FormValue("omega"), 64)
		tb, _ := strconv.ParseFloat(r.FormValue("tb"), 64)
		Pm, _ := strconv.ParseFloat(r.FormValue("Pm"), 64)
		Tm, _ := strconv.ParseFloat(r.FormValue("Tm"), 64)
		kp, _ := strconv.ParseFloat(r.FormValue("kp"), 64)
		zPerA, _ := strconv.ParseFloat(r.FormValue("zPerA"), 64)
		zPerP, _ := strconv.ParseFloat(r.FormValue("zPerP"), 64)

		result = calculateTask2(omega, tb, Pm, Tm, kp, zPerA, zPerP)
	}
	templates.ExecuteTemplate(w, "task2.html", result)
}

func calculateTask1(elements string, n float64) string {
	omegaMap := map[string]float64{"ПЛ-110 кВ": 0.07, "ПЛ-35 кВ": 0.02, "ПЛ-10 кВ": 0.02, "КЛ-10 кВ (траншея)": 0.03, "КЛ-10 кВ (кабельний канал)": 0.005, "Т-110 кВ": 0.015, "Т-35 кВ": 0.02, "Т-10 кВ (кабельна мережа 10 кВ)": 0.005, "Т-10 кВ (повітряна мережа 10 кВ)": 0.05}
	tvMap := map[string]float64{"ПЛ-110 кВ": 10.0, "ПЛ-35 кВ": 8.0, "ПЛ-10 кВ": 10.0, "КЛ-10 кВ (траншея)": 44.0, "КЛ-10 кВ (кабельний канал)": 17.5, "Т-110 кВ": 100.0, "Т-35 кВ": 80.0, "Т-10 кВ (кабельна мережа 10 кВ)": 60.0, "Т-10 кВ (повітряна мережа 10 кВ)": 60.0}
	tpMap := map[string]float64{"ПЛ-110 кВ": 35.0, "ПЛ-35 кВ": 35.0, "ПЛ-10 кВ": 35.0, "КЛ-10 кВ (траншея)": 9.0, "КЛ-10 кВ (кабельний канал)": 9.0, "Т-110 кВ": 43.0, "Т-35 кВ": 28.0, "Т-10 кВ (кабельна мережа 10 кВ)": 10.0, "Т-10 кВ (повітряна мережа 10 кВ)": 10.0}

	omegaSum, tRecovery, maxTp := 0.0, 0.0, 0.0
	for _, el := range strings.Split(elements, " ") {
		omegaSum += omegaMap[el]
		tRecovery += omegaMap[el] * tvMap[el]
		if tpMap[el] > maxTp {
			maxTp = tpMap[el]
		}
	}

	omegaSum += 0.03 * n
	tRecovery += 0.06 * n
	tRecovery /= omegaSum

	kAP := omegaSum * tRecovery / 8760
	kPP := 1.2 * maxTp / 8760
	omegaDK := 2 * omegaSum * (kAP + kPP)
	omegaDKS := omegaDK + 0.02

	return fmt.Sprintf(`Частота відмов: %.5f рік^-1\nСередня тривалість відновлення: %.5f год\nКоефіцієнт аварійного простою: %.5f\nКоефіцієнт планового простою: %.5f\nЧастота відмов двоколової системи: %.5f рік^-1\nЧастота відмов з секційним вимикачем: %.5f рік^-1`,
		omegaSum, tRecovery, kAP, kPP, omegaDK, omegaDKS)
}

func calculateTask2(omega, tb, Pm, Tm, kp, zPerA, zPerP float64) string {
	MWA := omega * tb * Pm * Tm
	MWP := kp * Pm * Tm
	M := zPerA*MWA + zPerP*MWP
	return fmt.Sprintf(`Аварійне недовідпущення: %.5f кВт·год\nПланове недовідпущення: %.5f кВт·год\nЗбитки: %.5f грн`, MWA, MWP, M)
}
