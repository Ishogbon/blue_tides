package database

import (
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client     *mongo.Client
	context    context.Context
	cancelFunc context.CancelFunc
	Host       string
	Port       uint16
	User       string
	Pass       string
}

func (db *Database) Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	db.cancelFunc = cancel
	db.context = ctx
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+db.User+":"+db.Pass+"@"+db.Host+":"+strconv.Itoa(int(db.Port))))
	if err != nil {
		panic(err)
	}
	db.client = client
	return client
}

func (db *Database) Close() {
	db.cancelFunc()
	if err := db.client.Disconnect(db.context); err != nil {
		panic(err)
	}
}

func (db *Database) InsertOne(dbName string, collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := db.client.Database(dbName).Collection(collectionName)
	result, err := collection.InsertOne(db.context, document)
	return result, err
}
