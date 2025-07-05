package utils

import (
	"bookcabin-seatmap-backend/models"
	"encoding/json"
	"os"
)

func LoadSegment(filePath string) (*models.Segment, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var parsed struct {
		SeatsItineraryParts []struct {
			SegmentSeatMaps []struct {
				Segment models.Segment `json:"segment"`
			} `json:"segmentSeatMaps"`
		} `json:"seatsItineraryParts"`
	}

	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, err
	}

	segment := parsed.
		SeatsItineraryParts[0].
		SegmentSeatMaps[0].
		Segment
	segment.ID = segment.Equipment + "-" + segment.Origin + "-" + segment.Destination

	return &segment, nil
}

func LoadPassenger(filePath string) (*models.Passenger, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var parsed struct {
		SeatsItineraryParts []struct {
			SegmentSeatMaps []struct {
				PassengerSeatMaps []struct {
					Passenger models.Passenger `json:"passenger"`
				} `json:"passengerSeatMaps"`
			} `json:"segmentSeatMaps"`
		} `json:"seatsItineraryParts"`
	}

	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, err
	}

	passenger := parsed.
		SeatsItineraryParts[0].
		SegmentSeatMaps[0].
		PassengerSeatMaps[0].
		Passenger
	passenger.ID = passenger.PassengerNumber

	return &passenger, nil
}

func LoadSeatMap(filePath string) (*models.SeatMap, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var parsed struct {
		SeatsItineraryParts []struct {
			SegmentSeatMaps []struct {
				PassengerSeatMaps []struct {
					SeatMap struct {
						Aircraft string         `json:"aircraft"`
						Cabins   []models.Cabin `json:"cabins"`
					} `json:"seatMap"`
				} `json:"passengerSeatMaps"`
			} `json:"segmentSeatMaps"`
		} `json:"seatsItineraryParts"`
	}

	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, err
	}

	seatMap := parsed.
		SeatsItineraryParts[0].
		SegmentSeatMaps[0].
		PassengerSeatMaps[0].
		SeatMap

	return &models.SeatMap{
		ID:       seatMap.Aircraft,
		Aircraft: seatMap.Aircraft,
		Cabins:   seatMap.Cabins,
	}, nil
}
