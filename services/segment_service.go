package services

import (
	"bookcabin-seatmap-backend/config"
	"bookcabin-seatmap-backend/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSegment(equipment string) (*models.Segment, error) {
	var segment models.Segment

	now := time.Now().Format(time.RFC3339)

	filter := bson.M{
		"equipment": bson.M{"$eq": equipment},
		"departure": bson.M{"$gt": now},
	}

	opts := options.FindOne().SetSort(bson.D{{Key: "departure", Value: 1}})

	err := config.SegmentCollection.FindOne(context.TODO(), filter, opts).Decode(&segment)
	if err != nil {
		return nil, err
	}

	return &segment, nil
}
