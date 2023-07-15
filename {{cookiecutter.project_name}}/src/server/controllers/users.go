package controllers

import (
	"context"
	"log"
	"mongo-test/api"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "users"

func (s Server) GetUsers(ctx context.Context, request api.GetUsersRequestObject) (api.GetUsersResponseObject, error) {
	results := api.Users{}
	filter := bson.M{}
	coll := s.DB.Collection(collectionName)
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return api.GetUsers200JSONResponse(results), nil
}

func (s Server) AddUsers(ctx context.Context, request api.AddUsersRequestObject) (api.AddUsersResponseObject, error) {
	// Create a slice of WriteModels to perform in bulk
	var writeModels []mongo.WriteModel
	for _, user := range *request.Body {
		model := mongo.NewUpdateOneModel()
		model.SetFilter(bson.M{"email": user.Email})
		model.SetUpdate(
			bson.M{
				"$set":         bson.M{"firstName": user.FirstName, "lastName": user.LastName},
				"$setOnInsert": bson.M{"email": user.Email},
			},
		)
		model.SetUpsert(true)
		writeModels = append(writeModels, model)
	}

	// If there are no write models, return a success response
	if len(writeModels) == 0 {
		usersAdded := api.UsersAdded{
			Status:       "no users to add",
			NumRequested: 0,
			NumAdded:     0,
			NumExisted:   0,
		}
		return api.AddUsers200JSONResponse(usersAdded), nil
	}

	// Perform the bulk write
	coll := s.DB.Collection(collectionName)
	opts := options.BulkWrite().SetOrdered(false)
	bulkWrite, err := coll.BulkWrite(context.TODO(), writeModels, opts)
	if err != nil {
		return api.AddUsers500Response{}, err
	}

	// Return a success response
	usersAdded := api.UsersAdded{
		Status:       "success",
		NumRequested: len(*request.Body),
		NumAdded:     int(bulkWrite.UpsertedCount),
		NumExisted:   int(bulkWrite.MatchedCount),
	}
	return api.AddUsers200JSONResponse(usersAdded), nil
}
