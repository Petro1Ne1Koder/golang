package main

import (
	"example.com/m/Services"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:8080")

	routerService := Services.NewRouterService()
	routerService.InitializeRoutes()

	err := http.ListenAndServe(":8080", routerService.GetHandler())
	if err != nil {
		log.Fatal(err)
	}
}
