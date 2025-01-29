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

	router.HandleFunc("/go/system/health", controllers.HealthHandler)
	router.HandleFunc("/go/system/version", controllers.VersionHandler)
	router.HandleFunc("/go/item", controllers.ItemHandler)
	router.HandleFunc("/go/item/{id}", controllers.ItemHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: middleware.Logging(router),
	}
	log.Println("Starting server on port " + port)
	server.ListenAndServe()
}
