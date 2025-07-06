package models

type SeatMap struct {
	ID                 string  `bson:"_id" json:"id"`
	RowsDisabledCauses []int   `json:"rowsDisabledCauses" bson:"rowsDisabledCauses"`
	Aircraft           string  `json:"aircraft" bson:"aircraft"`
	Cabins             []Cabin `json:"cabins" bson:"cabins"`
}

type Cabin struct {
	Deck        string   `json:"deck" bson:"deck"`
	SeatColumns []string `json:"seatColumns" bson:"seatColumns"`
	SeatRows    []Row    `json:"seatRows" bson:"seatRows"`
}

type Row struct {
	RowNumber int      `json:"rowNumber" bson:"rowNumber"`
	SeatCodes []string `json:"seatCodes" bson:"seatCodes"`
	Seats     []Seat   `json:"seats" bson:"seats"`
}

type Seat struct {
	Code           string `json:"code" bson:"code"`
	Available      bool   `json:"available" bson:"available"`
	FreeOfCharge   bool   `json:"freeOfCharge" bson:"freeOfCharge"`
	StorefrontCode string `json:"storefrontSlotCode" bson:"storefrontSlotCode"`
	Prices         struct {
		Alternatives [][]struct {
			Amount   float64 `json:"amount" bson:"amount"`
			Currency string  `json:"currency" bson:"currency"`
		} `json:"alternatives" bson:"alternatives"`
	} `json:"prices" bson:"prices"`
}
