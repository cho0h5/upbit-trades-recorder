package main

import (
        "context"

        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"
)

const URI = "mongodb://localhost:27017/"

type DB struct {
	client	*mongo.Client
	coll	*mongo.Collection
}

func ConnectDB() DB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))

        if err != nil {
                panic(err)
        }

	return DB { client, client.Database("test").Collection("books") }
}

func (db *DB) Disconnect() {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		panic(err)
        }
}

func (db *DB) InsertOne(doc interface{}) {
	_, err := db.coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
}
