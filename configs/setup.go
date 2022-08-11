package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CONNECT_DB() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(EnvironmentMongoUri()))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGO-DB")
	return client
}

var DB *mongo.Client = CONNECT_DB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangGinAPI").Collection(collectionName)
	return collection
}
