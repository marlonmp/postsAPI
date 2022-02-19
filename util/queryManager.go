package util

import (
	"context"
	"time"

	"github.com/marlonmp/postsAPI/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type queryManager struct {
	client *mongo.Client

	database *mongo.Database

	collection *mongo.Collection

	ctx    context.Context
	cancel context.CancelFunc

	err error
}

func NewQuery(databaseName, collectionName string) *queryManager {
	client, err := db.GetClient()

	if err != nil {
		return &queryManager{
			err: err,
		}
	}

	database := client.Database(databaseName)
	collection := database.Collection(collectionName)

	return &queryManager{
		client:     client,
		database:   database,
		collection: collection,
	}
}

func (Qm *queryManager) createContext() {
	Qm.ctx, Qm.cancel = context.WithTimeout(context.Background(), 5*time.Second)
}

func (Qm *queryManager) FindOne(filter interface{}) (*mongo.SingleResult, error) {

	if Qm.err != nil {
		return nil, Qm.err
	}

	Qm.createContext()

	defer Qm.cancel()

	return Qm.collection.FindOne(Qm.ctx, filter), nil
}

func (Qm *queryManager) Find(filter interface{}) (*mongo.Cursor, error) {

	if Qm.err != nil {
		return nil, Qm.err
	}

	Qm.createContext()

	defer Qm.cancel()

	return Qm.collection.Find(Qm.ctx, filter)
}

func (Qm *queryManager) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {

	if Qm.err != nil {
		return nil, Qm.err
	}

	Qm.createContext()

	defer Qm.cancel()

	return Qm.collection.InsertOne(Qm.ctx, document)
}

func (Qm *queryManager) UpdateOne(filter, update interface{}) (*mongo.UpdateResult, error) {

	if Qm.err != nil {
		return nil, Qm.err
	}

	Qm.createContext()

	defer Qm.cancel()

	return Qm.collection.UpdateOne(Qm.ctx, filter, update)
}

func (Qm *queryManager) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {

	if Qm.err != nil {
		return nil, Qm.err
	}

	Qm.createContext()

	defer Qm.cancel()

	return Qm.collection.DeleteOne(Qm.ctx, filter)
}
