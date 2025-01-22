package database

import (
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client  *mongo.Client
	context context.Context
	Host    string
	Port    uint16
	User    string
	Pass    string
}

func (db *Database) Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	db.context = ctx
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+db.User+":"+db.Pass+"@"+db.Host+":"+strconv.Itoa(int(db.Port))))
	if err != nil {
		panic(err)
	}
	db.client = client
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return client
}
