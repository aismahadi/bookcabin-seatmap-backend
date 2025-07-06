package services

import (
	"bookcabin-seatmap-backend/config"
	"bookcabin-seatmap-backend/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPassenger(id string) (*models.Passenger, error) {
	var passenger models.Passenger

	err := config.PassengerCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&passenger)
	if err != nil {
		return nil, err
	}

	return &passenger, nil
}
