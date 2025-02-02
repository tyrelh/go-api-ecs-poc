package main

import (
	"go-api-poc/api"
	"go-api-poc/controllers"
	"go-api-poc/db"
	"go-api-poc/middleware"
	"log"
	"net/http"
)

var (
	port = "8080"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	strictServerDefinition := api.NewStrictHandler(
		controllers.NewStrictServer(),
		[]api.StrictMiddlewareFunc{},
	)
	httpHandler := api.HandlerWithOptions(strictServerDefinition, api.StdHTTPServerOptions{
		BaseRouter: http.NewServeMux(),
		Middlewares: []api.MiddlewareFunc{
			middleware.Logging,
		},
	})
	server := &http.Server{
		Handler: httpHandler,
		Addr:    ":" + port,
	}

	err := db.ConnectToDb()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Starting server on port " + port + "...")
	log.Fatal(server.ListenAndServe())
}
