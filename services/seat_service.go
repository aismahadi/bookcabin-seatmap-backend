package services

import (
	"bookcabin-seatmap-backend/config"
	"bookcabin-seatmap-backend/models"
	"context"
)

func GetFullSeatMap() *models.Aircraft {
	var result models.Aircraft
	config.SeatMapCollection.FindOne(context.TODO(), map[string]interface{}{}).Decode(&result)
	return &result
}
