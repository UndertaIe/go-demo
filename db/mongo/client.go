package mongoclient

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBURI = "mongodb+srv://604712799:mongo007@cluster0.iwvcfsl.mongodb.net/?retryWrites=true&w=majority"

// mongoexport --uri mongodb+srv://604712799:mongo007@cluster0.iwvcfsl.mongodb.net/sample_airbnb --collection listingsAndReviews --type json --out sample01.json

func MongoClient() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MongoDBURI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	return client
}
