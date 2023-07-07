package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	DB *mongo.Database
}
