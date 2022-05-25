package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jwt-auth.com/config"
	"log"
	"time"
)

func DBinstance() *mongo.Client {
	MongoDB := config.Getconfig().MongoUrl
	// fmt.Print(MongoDB)
	// credential := options.Credential{
	// 	Username: "nafiul1916",
	// 	Password: "quddus1916",
	// }

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("JWT").Collection(collectionName)
	return collection
}
