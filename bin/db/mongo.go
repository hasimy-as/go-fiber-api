package db

import (
	"context"
	"log"
	"time"

	"github.com/hasimy-as/go-fiber-api/bin/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GlobalEnv.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	cx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(cx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(cx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(config.GlobalEnv.DBName).Collection(collectionName)
	return collection
}
