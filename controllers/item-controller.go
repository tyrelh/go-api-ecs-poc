package controllers

import (
	"net/http"

	"go-api-poc/services"
)

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	switch r.Method {
	case http.MethodGet:
		if id == "" {
			services.GetItems(w, r)
		}
		services.GetItem(w, r, id)
	case http.MethodPost:
		services.CreateItem(w, r)
	case http.MethodPut:
		services.UpdateItem(w, r, id)
	case http.MethodDelete:
		services.DeleteItem(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
