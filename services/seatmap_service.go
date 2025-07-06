package services

import (
	"bookcabin-seatmap-backend/config"
	"bookcabin-seatmap-backend/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSeatMap(id string) (*models.SeatMap, error) {
	var seatmap models.SeatMap

	err := config.SeatMapCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&seatmap)
	if err != nil {
		return nil, err
	}

	return &seatmap, nil
}
