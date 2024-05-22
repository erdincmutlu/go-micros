package main

import (
	"context"
	"log"
	"log-service/data"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogItem(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})
	if err != nil {
		log.Println("Error writing to Mongo: ", err.Error())

		return err
	}

	*resp = "Processed payload via PRC: " + payload.Name
	return nil
}
