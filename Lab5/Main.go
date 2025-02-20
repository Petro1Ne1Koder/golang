package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/task1", Task1Handler)
	http.HandleFunc("/task2", Task2Handler)

	fmt.Println("Сервер запущено на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}