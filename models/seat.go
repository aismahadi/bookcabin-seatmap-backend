package models

type Seat struct {
	Code           string `json:"code" bson:"code"` // e.g. "12A"
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
