package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
}

func main() {
	// connect to mongo
	ctx := context.Background()
	mongoClient, err := connectToMongo(ctx)
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}()
}

func connectToMongo(ctx context.Context) (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Error connecting: ", err.Error())
		return nil, err
	}

	return c, nil
}
