package configuration

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDataBase() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetMongoUrl()))
	if err != nil {
		fmt.Println("Error connect to Data base ...")

		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		fmt.Println("Error connect to Data base ...")

		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Data base ...")
	return client
}

var DbClient *mongo.Client = ConnectDataBase()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("go-gin-api").Collection(collectionName)
	return collection
}
