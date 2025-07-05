package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var SeatMapCollection *mongo.Collection

func ConnectMongo() {
	username := "mongo"
	password := "mongo"
	host := "localhost"
	port := 27017
	dbName := "bookcabin"

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	MongoClient = client
	SeatMapCollection = client.Database(dbName).Collection("seatmaps")
	log.Println("Connected to MongoDB with auth")
}
