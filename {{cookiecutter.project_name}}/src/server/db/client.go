package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(uri string, username string, password string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	clientOptions := options.Client()
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})
	clientOptions.SetRegistry(mongoRegistry)
	clientOptions.ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	return client, err
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
