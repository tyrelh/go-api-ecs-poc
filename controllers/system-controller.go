package controllers

import (
	"go-api-poc/services"
	"net/http"

	"go-api-poc/api"
)

func (Server) GetGoSystemVersion(w http.ResponseWriter, r *http.Request) {
	resp := api.GetGoSystemVersion200JSONResponse{
		Version: services.GetVersion(),
	}
	resp.VisitGetGoSystemVersionResponse(w)

	// w.WriteHeader(http.StatusOK)
	// _ = json.NewEncoder(w).Encode(resp)
}

func (Server) GetGoSystemHealth(w http.ResponseWriter, r *http.Request) {
	health := "I'm ok."
	resp := api.GetGoSystemHealth200JSONResponse{
		Status: &health,
	}
	resp.VisitGetGoSystemHealthResponse(w)

	// w.WriteHeader(http.StatusOK)
	// _ = json.NewEncoder(w).Encode(resp)
}
