package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logger").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:     payload.Name,
		Data:     payload.Data,
		CreateAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "processed payload via rpc:" + payload.Name
	return nil
}
