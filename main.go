package main

import (
	"go-api-poc/api"
	"go-api-poc/controllers"
	"go-api-poc/middleware"
	"log"
	"net/http"
	// "go-api-poc/controllers"
	// "go-api-poc/middleware"
)

var (
	port = "8080"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	serverDefinition := controllers.NewServer()
	router := http.NewServeMux()
	httpHandler := api.HandlerWithOptions(serverDefinition, api.StdHTTPServerOptions{
		BaseRouter: router,
		Middlewares: []api.MiddlewareFunc{
			middleware.Logging,
		},
	})
	server := &http.Server{
		Handler: httpHandler,
		Addr:    ":" + port,
	}
	log.Println("Starting server on port " + port + "...")
	log.Fatal(server.ListenAndServe())
}
