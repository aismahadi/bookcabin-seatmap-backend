package utils

import (
	"bookcabin-seatmap-backend/models"
	"encoding/json"
	"os"
)

func LoadData(filePath string) ([]*models.Segment, []*models.Passenger, []*models.SeatMap, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, nil, err
	}

	var parsed struct {
		SeatsItineraryParts []struct {
			SegmentSeatMaps []struct {
				Segment           models.Segment `json:"segment"`
				PassengerSeatMaps []struct {
					Passenger models.Passenger `json:"passenger"`
					SeatMap   struct {
						Aircraft string         `json:"aircraft"`
						Cabins   []models.Cabin `json:"cabins"`
					} `json:"seatMap"`
				} `json:"passengerSeatMaps"`
			} `json:"segmentSeatMaps"`
		} `json:"seatsItineraryParts"`
	}

	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, nil, nil, err
	}

	var segments []*models.Segment
	var passengers []*models.Passenger
	var seatmaps []*models.SeatMap

	for _, itinerary := range parsed.SeatsItineraryParts {
		for _, segmentMap := range itinerary.SegmentSeatMaps {
			// Add segment
			segment := segmentMap.Segment
			segment.ID = segment.SegmentRef
			segments = append(segments, &segment)

			for _, psm := range segmentMap.PassengerSeatMaps {
				// Add passenger
				passenger := psm.Passenger
				passenger.ID = passenger.PassengerNumber
				passengers = append(passengers, &passenger)

				// Add seatmap
				seatMap := &models.SeatMap{
					ID:       psm.SeatMap.Aircraft,
					Aircraft: psm.SeatMap.Aircraft,
					Cabins:   psm.SeatMap.Cabins,
				}
				seatmaps = append(seatmaps, seatMap)
			}
		}
	}

	return segments, passengers, seatmaps, nil
}
