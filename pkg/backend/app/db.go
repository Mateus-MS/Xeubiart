package app

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartDBConnection() (mongoClient *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	if mongoClient, err = mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://adm:adm@192.168.1.94:5432")); err != nil {
		log.Fatal("Mongo connection error: " + err.Error())
	}

	return mongoClient
}
