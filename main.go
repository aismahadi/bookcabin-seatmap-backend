package main

import (
	"bookcabin-seatmap-backend/config"
	"bookcabin-seatmap-backend/routes"
	"bookcabin-seatmap-backend/utils"
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	config.ConnectMongo()

	// Load all data
	segments, passengers, seatMaps, err := utils.LoadData("SeatMapResponse.json")
	if err != nil {
		log.Fatal("Failed to load data:", err)
	}

	// Clear collections before insert
	config.SegmentCollection.DeleteMany(context.TODO(), bson.M{})
	config.PassengerCollection.DeleteMany(context.TODO(), bson.M{})
	config.SeatMapCollection.DeleteMany(context.TODO(), bson.M{})

	// Insert segments
	segmentDocs := make([]interface{}, len(segments))
	for i, s := range segments {
		segmentDocs[i] = s
	}
	if len(segmentDocs) > 0 {
		_, err = config.SegmentCollection.InsertMany(context.TODO(), segmentDocs)
		if err != nil {
			log.Fatal("Failed to insert segments:", err)
		}
		log.Println("Segments seeded into MongoDB")
	}

	// Insert passengers
	passengerDocs := make([]interface{}, len(passengers))
	for i, p := range passengers {
		passengerDocs[i] = p
	}
	if len(passengerDocs) > 0 {
		_, err = config.PassengerCollection.InsertMany(context.TODO(), passengerDocs)
		if err != nil {
			log.Fatal("Failed to insert passengers:", err)
		}
		log.Println("Passengers seeded into MongoDB")
	}

	// Insert seat maps
	seatMapDocs := make([]interface{}, len(seatMaps))
	for i, sm := range seatMaps {
		seatMapDocs[i] = sm
	}
	if len(seatMapDocs) > 0 {
		_, err = config.SeatMapCollection.InsertMany(context.TODO(), seatMapDocs)
		if err != nil {
			log.Fatal("Failed to insert seatmaps:", err)
		}
		log.Println("Seat maps seeded into MongoDB")
	}

	// Start server
	r := gin.Default()
	r.Use(cors.Default())

	routes.RegisterSegmentRoutes(r)
	routes.RegisterSeatMapRoutes(r)
	routes.RegisterPassengerRoutes(r)

	r.Run(":8080")
}
