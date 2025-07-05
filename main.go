package main

import (
	"bookcabin-seatmap-backend/config"
	"bookcabin-seatmap-backend/routes"
	"bookcabin-seatmap-backend/utils"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.ConnectMongo()

	aircraft, err := utils.LoadSeatMap("SeatMapResponse.json")
	if err != nil {
		log.Fatal(err)
	}

	config.SeatMapCollection.DeleteMany(context.TODO(), map[string]interface{}{})

	_, err = config.SeatMapCollection.InsertOne(context.TODO(), aircraft)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data seeded into MongoDB")

	r := gin.Default()
	r.Use(cors.Default())
	routes.RegisterSeatRoutes(r)
	r.Run(":8080")
}
