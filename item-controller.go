package main

import (
	"fmt"
	"net/http"
)

// Create a custom response writer to capture the output
type responseWriter struct {
	http.ResponseWriter
	body []byte
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)
	return rw.ResponseWriter.Write(b)
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	// Wrap the original ResponseWriter
	customWriter := &responseWriter{ResponseWriter: w}

	switch r.Method {
	case http.MethodGet:
		getItems(customWriter, r)
		// Log or print the response
		fmt.Printf("Response body: %s\n", string(customWriter.body))
	case http.MethodPost:
		createItem(customWriter, r)
		fmt.Printf("Response body: %s\n", string(customWriter.body))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	switch r.Method {
	case http.MethodGet:
		getItem(w, r, id)
	case http.MethodPut:
		updateItem(w, r, id)
	case http.MethodDelete:
		deleteItem(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
