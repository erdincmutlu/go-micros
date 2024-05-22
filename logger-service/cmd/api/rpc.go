package main

import (
	"context"
	"log"
	"log-service/data"
)

// RPCServer is the type for our RPC Server. Methods that take this as a
// receiver are available over RPC, as long as they are exported.
type RPCServer struct {
}

// RPCPayload is the type for data we receive from RPC
type RPCPayload struct {
	Name string
	Data string
}

// LogInfo writes our payload to mongo
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})
	if err != nil {
		log.Println("Error writing to Mongo: ", err.Error())

		return err
	}

	// resp is the message sent back to the RPC caller
	*resp = "Processed payload via PRC: " + payload.Name
	return nil
}
