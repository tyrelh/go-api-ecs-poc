package main

import (
	"log"
	"net/http"
	"sync"

	"go-api-poc/middleware"
)

var (
	items = make(map[string]Item)
	mu    sync.Mutex
	port  = "8080"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/items/{id}", itemHandler)
	router.HandleFunc("/items", itemsHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: middleware.Logging(router),
	}
	log.Println("Starting server on port " + port)
	server.ListenAndServe()
}
