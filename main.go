package main

import (
	"log"
	"net/http"

	"go-api-poc/controllers"
	"go-api-poc/middleware"
)

var (
	port = "8080"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/go/items/{id}", controllers.ItemHandler)
	router.HandleFunc("/go/items", controllers.ItemsHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: middleware.Logging(router),
	}
	log.Println("Starting server on port " + port)
	server.ListenAndServe()
}
