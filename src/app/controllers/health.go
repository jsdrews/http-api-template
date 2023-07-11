package controllers

import (
	"context"
	"fmt"

	"mongo-test/api"

	"go.mongodb.org/mongo-driver/bson"
)

func (s Server) GetHealth(ctx context.Context, request api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	var result bson.M
	err := s.DB.RunCommand(context.TODO(), bson.M{"dbStats": 1}).Decode(&result)
	if err != nil {
		return api.GetHealth500Response{}, err
	}
	
	// Check if the database is healthy
	if result["ok"] != 1.0 {
		return api.GetHealth503Response{}, fmt.Errorf("database is not healthy")
	}

	// Report back success
	return api.GetHealth200JSONResponse(
		api.HealthCheck{
			Status: "OK",
		},
	), nil
}
