package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var SeatMapCollection *mongo.Collection
var PassengerCollection *mongo.Collection
var SegmentCollection *mongo.Collection

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

	db := client.Database(dbName)

	// 🔍 List current collections
	collections, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatal("Failed to list collections:", err)
	}

	// Create "seatmaps" collection if not exists
	if !contains(collections, "seatmaps") {
		err := db.CreateCollection(ctx, "seatmaps")
		if err != nil {
			log.Fatal("Failed to create 'seatmaps' collection:", err)
		}
		log.Println("Created 'seatmaps' collection")
	}

	// Create "passengers" collection if not exists
	if !contains(collections, "passengers") {
		err := db.CreateCollection(ctx, "passengers")
		if err != nil {
			log.Fatal("Failed to create 'passengers' collection:", err)
		}
		log.Println("Created 'passengers' collection")
	}

	// Create "segments" collection if not exists
	if !contains(collections, "segments") {
		err := db.CreateCollection(ctx, "segments")
		if err != nil {
			log.Fatal("Failed to create 'segments' collection:", err)
		}
		log.Println("Created 'segments' collection")
	}

	// Assign global collections
	MongoClient = client
	SeatMapCollection = db.Collection("seatmaps")
	PassengerCollection = db.Collection("passengers")
	SegmentCollection = db.Collection("segments")

	log.Println("Connected to MongoDB with auth")
}

// Helper function
func contains(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}
