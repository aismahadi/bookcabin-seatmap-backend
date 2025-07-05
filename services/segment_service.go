package services

import (
	"bookcabin-seatmap-backend/config"
	"bookcabin-seatmap-backend/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSegment() (*models.Segment, error) {
	var segment models.Segment

	err := config.SegmentCollection.FindOne(context.TODO(), bson.M{}).Decode(&segment)
	if err != nil {
		return nil, err
	}

	return &segment, nil
}
