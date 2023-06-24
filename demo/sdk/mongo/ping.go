package mongo

import (
	"context"
	"fmt"

	mongoclient "github.com/UndertaIe/go-demo/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func (Mongo) Ping() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	client := mongoclient.MongoClient()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

