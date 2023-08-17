package controllers

import (
	"context"
	"time"

	"server-api/api"

	"go.mongodb.org/mongo-driver/bson"
)

func (s Server) GetHealth(ctx context.Context, request api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	var result bson.M
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err := s.DB.RunCommand(ctx, bson.M{"dbStats": 1}).Decode(&result)
	if err != nil {
		return api.GetHealth500JSONResponse{
			Error: err.Error(),
		}, nil
	}

	// Check if the database is healthy
	_, found := result["ok"]
	if !found {
		s.Logger.Error("database is not healthy: %v", result)
		return api.GetHealth503JSONResponse{
			Error: "database is not healthy",
		}, nil
	}

	// Report back success
	return api.GetHealth200JSONResponse(
		api.HealthCheck{
			Status: "OK",
		},
	), nil
}
