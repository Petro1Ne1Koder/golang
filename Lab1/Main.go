package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Структура топлива
type FuelComposition struct {
	H, C, S, N, O, W, A, Q float64
}

// Структура для мазуту
type MazutComposition struct {
	C, H, O, S, V, W, A, Qdaf float64
}

// Функція розрахунку сухої та паливної маси
func calculateComposition(f FuelComposition) (FuelComposition, FuelComposition, float64, float64, float64, float64, float64) {
	KRS := 100 / (100 - f.W)
	KRG := 100 / (100 - f.W - f.A)

	sDry := FuelComposition{
		H: f.H * KRS,
		C: f.C * KRS,
		S: f.S * KRS,
		N: f.N * KRS,
		O: f.O * KRS,
		A: f.A * KRS,
	}

	sCombustible := FuelComposition{
		H: f.H * KRG,
		C: f.C * KRG,
		S: f.S * KRG,
		N: f.N * KRG,
		O: f.O * KRG,
	}

	QpH := (339*f.C + 1030*f.H - 108.8*(f.O-f.S) - 25*f.W) / 1000
	QdH := QpH * KRS
	QdafH := QpH * KRG

	return sDry, sCombustible, QpH, QdH, QdafH, KRS, KRG
}

// Функція розрахунку робочої маси мазуту
func calculateMazutComposition(f MazutComposition) (MazutComposition, float64) {
	Kp := (100 - f.W - f.A) / 100

	mP := MazutComposition{
		C:    f.C * Kp,
		H:    f.H * Kp,
		O:    f.O * Kp,
		S:    f.S * Kp,
		V:    f.V * Kp,
		W:    f.W,
		A:    f.A,
		Qdaf: f.Qdaf,
	}

	Qp := f.Qdaf*(100-f.W-f.A)/100 - 0.025*f.W
	return mP, Qp
}

// Обробник сторінки для завдання 1
func handlerFuel(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	var f FuelComposition
	f.H, _ = strconv.ParseFloat(r.FormValue("H"), 64)
	f.C, _ = strconv.ParseFloat(r.FormValue("C"), 64)
	f.S, _ = strconv.ParseFloat(r.FormValue("S"), 64)
	f.N, _ = strconv.ParseFloat(r.FormValue("N"), 64)
	f.O, _ = strconv.ParseFloat(r.FormValue("O"), 64)
	f.W, _ = strconv.ParseFloat(r.FormValue("W"), 64)
	f.A, _ = strconv.ParseFloat(r.FormValue("A"), 64)

	sDry, sCombustible, QpH, QdH, QdafH, KRS, KRG := calculateComposition(f)

	data := struct {
		Input, DryComposition, CombustibleComposition FuelComposition
		QpH, QdH, QdafH, KRS, KRG                     float64
	}{
		Input:                  f,
		DryComposition:         sDry,
		CombustibleComposition: sCombustible,
		QpH:                    QpH,
		QdH:                    QdH,
		QdafH:                  QdafH,
		KRS:                    KRS,
		KRG:                    KRG,
	}

	tmpl.Execute(w, data)
}

// Обробник сторінки для завдання 2
func handlerMazut(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/mobile.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	var f MazutComposition
	f.C, _ = strconv.ParseFloat(r.FormValue("C"), 64)
	f.H, _ = strconv.ParseFloat(r.FormValue("H"), 64)
	f.O, _ = strconv.ParseFloat(r.FormValue("O"), 64)
	f.S, _ = strconv.ParseFloat(r.FormValue("S"), 64)
	f.V, _ = strconv.ParseFloat(r.FormValue("V"), 64)
	f.W, _ = strconv.ParseFloat(r.FormValue("W"), 64)
	f.A, _ = strconv.ParseFloat(r.FormValue("A"), 64)
	f.Qdaf, _ = strconv.ParseFloat(r.FormValue("Qdaf"), 64)

	mP, Qp := calculateMazutComposition(f)

	data := struct {
		Input  MazutComposition
		Output MazutComposition
		Qp     float64
	}{
		Input:  f,
		Output: mP,
		Qp:     Qp,
	}

	tmpl.Execute(w, data)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlerFuel)
	http.HandleFunc("/mazut", handlerMazut)
	fmt.Println("Server is running on http://localhost:8080 and http://localhost:8080/mazut")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
