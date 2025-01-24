package controllers

import (
	"encoding/json"
	"go-api-poc/services"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := healthResponse{Status: "ItsGood"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	currentVersion := services.GetVersion()
	response := versionResponse{Version: currentVersion}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type healthResponse struct {
	Status string `json:"status"`
}

type versionResponse struct {
	Version string `json:"version"`
}
