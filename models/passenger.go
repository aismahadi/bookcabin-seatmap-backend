package models

type Passenger struct {
	ID               string           `bson:"_id" json:"id"`
	PassengerIndex   int              `bson:"passengerIndex" json:"passengerIndex"`
	PassengerNumber  string           `bson:"passengerNumber" json:"passengerNumber"`
	PassengerDetails PassengerDetails `bson:"passengerDetails" json:"passengerDetails"`
	PassengerInfo    PassengerInfo    `bson:"passengerInfo" json:"passengerInfo"`
	Preferences      Preferences      `bson:"preferences" json:"preferences"`
	DocumentInfo     DocumentInfo     `bson:"documentInfo" json:"documentInfo"`
}

type PassengerDetails struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}

type PassengerInfo struct {
	DateOfBirth string   `bson:"dateOfBirth" json:"dateOfBirth"`
	Gender      string   `bson:"gender" json:"gender"`
	Type        string   `bson:"type" json:"type"`
	Emails      []string `bson:"emails" json:"emails"`
	Phones      []string `bson:"phones" json:"phones"`
	Address     Address  `bson:"address" json:"address"`
}

type Address struct {
	Street1     string `bson:"street1" json:"street1"`
	Street2     string `bson:"street2" json:"street2"`
	Postcode    string `bson:"postcode" json:"postcode"`
	State       string `bson:"state" json:"state"`
	City        string `bson:"city" json:"city"`
	Country     string `bson:"country" json:"country"`
	AddressType string `bson:"addressType" json:"addressType"`
}

type Preferences struct {
	SpecialPreferences SpecialPreferences `bson:"specialPreferences" json:"specialPreferences"`
	FrequentFlyer      []FrequentFlyer    `bson:"frequentFlyer" json:"frequentFlyer"`
}

type SpecialPreferences struct {
	MealPreference               string   `bson:"mealPreference" json:"mealPreference"`
	SeatPreference               string   `bson:"seatPreference" json:"seatPreference"`
	SpecialRequests              []string `bson:"specialRequests" json:"specialRequests"`
	SpecialServiceRequestRemarks []string `bson:"specialServiceRequestRemarks" json:"specialServiceRequestRemarks"`
}

type FrequentFlyer struct {
	Airline    string `bson:"airline" json:"airline"`
	Number     string `bson:"number" json:"number"`
	TierNumber int    `bson:"tierNumber" json:"tierNumber"`
}

type DocumentInfo struct {
	IssuingCountry string `bson:"issuingCountry" json:"issuingCountry"`
	CountryOfBirth string `bson:"countryOfBirth" json:"countryOfBirth"`
	DocumentType   string `bson:"documentType" json:"documentType"`
	Nationality    string `bson:"nationality" json:"nationality"`
}
