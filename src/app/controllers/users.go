package controllers

import (
	"context"
	"mongo-test/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Server) GetUsers(ctx context.Context, request api.GetUsersRequestObject) (api.GetUsersResponseObject, error) {
	var result bson.M
	err := s.DB.RunCommand(context.TODO(), bson.M{"dbStats": 1}).Decode(&result)
	if err != nil {
		panic(err)
	}
	// c.JSON(http.StatusOK, result)
	// return api.GetUsers200JSONResponse(result), nil
	return api.GetUsersResponseObject(api.GetUsersdefaultJSONResponse{Body: result, StatusCode: http.StatusOK}), nil
}
