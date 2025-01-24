package controllers

import (
	"fmt"
	"net/http"

	"go-api-poc/services"
)

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	// Wrap the original ResponseWriter
	customWriter := &services.ResWriter{ResponseWriter: w}

	switch r.Method {
	case http.MethodGet:
		services.GetItems(customWriter, r)
	case http.MethodPost:
		services.CreateItem(customWriter, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	fmt.Printf("Response body: %s\n", string(customWriter.Body))
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	switch r.Method {
	case http.MethodGet:
		services.GetItem(w, r, id)
	case http.MethodPut:
		services.UpdateItem(w, r, id)
	case http.MethodDelete:
		services.DeleteItem(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
