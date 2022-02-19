package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientOptions *options.ClientOptions

func init() {
	clientOptions = options.Client()

	clientOptions.ApplyURI("mongodb://localhost:27017")
}

func GetClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	return mongo.Connect(ctx, clientOptions)
}
