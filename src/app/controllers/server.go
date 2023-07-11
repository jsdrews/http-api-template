package controllers

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	DB *mongo.Database
}

// Validate the body of the request using
// the go-playground/validator/v10 package
// https://pkg.go.dev/github.com/go-playground/validator/v10
func Validate(i interface{}) error {
	return validator.New().Struct(i)
}
