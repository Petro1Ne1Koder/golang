package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Подключение статических файлов
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Обработчики маршрутов
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/single-phase", SinglePhaseHandler)
	http.HandleFunc("/three-phase", ThreePhaseHandler)
	http.HandleFunc("/stability", StabilityHandler)

	fmt.Println("Сервер запущено на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


