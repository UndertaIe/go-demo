package mongo

import (
	"context"
	"encoding/json"
	"fmt"

	mongoclient "github.com/UndertaIe/go-demo/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (Mongo) FindOne() {
	client := mongoclient.MongoClient()
	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func (Mongo) FindOne2() {

	client := mongoclient.MongoClient()
	coll := client.Database("sample_restaurants").Collection("restaurants")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"rating", 5}}).Decode(&result)
	if err != nil {
		panic(err)
	}

}
