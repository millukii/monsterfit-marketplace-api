package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func NewMongoConn(uri string ,ctx 	context.Context)  (*mongo.Client, error){
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
  if err != nil {
    log.Print(err)
			return nil, err
  }

	return client, nil
}

func GetCollection(client *mongo.Client,db string, collectionName string, ctx context.Context) (*mongo.Collection, error){
	collection := client.Database(db).Collection(collectionName)

	if collection == nil{
		return nil, fmt.Errorf("collection not found")
	}
	return collection, nil
}