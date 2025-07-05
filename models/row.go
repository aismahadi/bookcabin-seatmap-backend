package models

type Row struct {
	RowNumber int      `json:"rowNumber" bson:"rowNumber"`
	SeatCodes []string `json:"seatCodes" bson:"seatCodes"` // Optional if needed
	Seats     []Seat   `json:"seats" bson:"seats"`
}
