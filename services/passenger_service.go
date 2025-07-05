package services

import (
	"bookcabin-seatmap-backend/config"
	"bookcabin-seatmap-backend/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPassenger() (*models.Passenger, error) {
	var passenger models.Passenger

	// You can customize the query if needed, for now fetch the first passenger
	err := config.PassengerCollection.FindOne(context.TODO(), bson.M{}).Decode(&passenger)
	if err != nil {
		return nil, err
	}

	return &passenger, nil
}
