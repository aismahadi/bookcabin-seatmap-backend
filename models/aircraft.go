package models

type Aircraft struct {
	ID       string  `bson:"_id" json:"id"`
	Aircraft string  `json:"aircraft" bson:"aircraft"`
	Cabins   []Cabin `json:"cabins" bson:"cabins"`
}

type Cabin struct {
	Deck        string   `json:"deck" bson:"deck"`
	SeatColumns []string `json:"seatColumns" bson:"seatColumns"`
	SeatRows    []Row    `json:"seatRows" bson:"seatRows"`
}
