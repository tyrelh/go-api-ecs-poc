package controllers

import (
	"context"
	"go-api-poc/services"

	"go-api-poc/api"
)

func (StrictServer) GetGoSystemVersion(ctx context.Context, request api.GetGoSystemVersionRequestObject) (api.GetGoSystemVersionResponseObject, error) {
	version := services.GetVersion()
	return api.GetGoSystemVersion200JSONResponse{
		Version: version,
	}, nil
}

func (StrictServer) GetGoSystemHealth(ctx context.Context, request api.GetGoSystemHealthRequestObject) (api.GetGoSystemHealthResponseObject, error) {
	health := "I'm ok."
	return api.GetGoSystemHealth200JSONResponse{
		Status: &health,
	}, nil
}
