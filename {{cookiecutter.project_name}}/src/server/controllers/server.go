package controllers

import (
	"context"
	"fmt"
	"log"
	"server-api/api"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

// AppLogger is a struct that contains
// pointers to the loggers that will be used
type AppLogger struct {
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
}

func (l AppLogger) Warn(format string, v ...interface{}) {
	l.WarningLogger.Printf(format, v...)
}

func (l AppLogger) Info(format string, v ...interface{}) {
	l.InfoLogger.Printf(format, v...)
}

func (l AppLogger) Error(format string, v ...interface{}) {
	l.ErrorLogger.Printf(format, v...)
}

// Server is a struct that contains
// pointers to the database and loggers
type Server struct {
	ApiVersion string
	DB         *mongo.Database
	Logger     *AppLogger
}

func (s Server) GetVersion(ctx context.Context, request api.GetVersionRequestObject) (api.GetVersionResponseObject, error) {
	return api.GetVersion200JSONResponse{
		Version: s.ApiVersion,
	}, nil
}

// Validate the body of the request using
// the go-playground/validator/v10 package
// https://pkg.go.dev/github.com/go-playground/validator/v10
func Validate(i interface{}) error {
	return validator.New().Struct(i)
}

// GetGooglePubSubMessageAttribute gets the value of the
// attribute with the given key from the Google Pub/Sub
// push subscription message
func GetGooglePubSubMessageAttribute(attributes *map[string]interface{}, key string) (string, error) {

	value, ok := (*attributes)[key]
	if !ok {
		return "", fmt.Errorf("cloud event missing attribute: %s", key)
	}
	convertedValue, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("cloud event attribute %s not a string", key)
	}
	return convertedValue, nil
}

// DecodeGooglePubSubMessage decodes the message
// from the Google Pub/Sub push subscription
// func DecodeGooglePubSubMessage(message *api.GooglePubSubPushMessage, decodeTo any) error {

// 	// decode message
// 	data, err := base64.StdEncoding.DecodeString(*message.Data)
// 	if err != nil {
// 		return err
// 	}

// 	// unmarshal data into decodeTo
// 	err = json.Unmarshal(data, &decodeTo)
// 	if err != nil {
// 		return err
// 	}

// 	// validate decodeTo
// 	err = Validate(decodeTo)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
