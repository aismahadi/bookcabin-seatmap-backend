package models

type Segment struct {
	ID                          string                  `bson:"_id" json:"id"`
	SegmentOfferInformation     SegmentOfferInformation `bson:"segmentOfferInformation" json:"segmentOfferInformation"`
	Duration                    int                     `bson:"duration" json:"duration"`
	CabinClass                  string                  `bson:"cabinClass" json:"cabinClass"`
	Equipment                   string                  `bson:"equipment" json:"equipment"`
	Flight                      FlightDetails           `bson:"flight" json:"flight"`
	Origin                      string                  `bson:"origin" json:"origin"`
	Destination                 string                  `bson:"destination" json:"destination"`
	Departure                   string                  `bson:"departure" json:"departure"` // ISO 8601
	Arrival                     string                  `bson:"arrival" json:"arrival"`     // ISO 8601
	BookingClass                string                  `bson:"bookingClass" json:"bookingClass"`
	LayoverDuration             int                     `bson:"layoverDuration" json:"layoverDuration"`
	FareBasis                   string                  `bson:"fareBasis" json:"fareBasis"`
	SubjectToGovernmentApproval bool                    `bson:"subjectToGovernmentApproval" json:"subjectToGovernmentApproval"`
	SegmentRef                  string                  `bson:"segmentRef" json:"segmentRef"`
}

type SegmentOfferInformation struct {
	FlightsMiles int  `bson:"flightsMiles" json:"flightsMiles"`
	AwardFare    bool `bson:"awardFare" json:"awardFare"`
}

type FlightDetails struct {
	FlightNumber          int      `bson:"flightNumber" json:"flightNumber"`
	OperatingFlightNumber int      `bson:"operatingFlightNumber" json:"operatingFlightNumber"`
	AirlineCode           string   `bson:"airlineCode" json:"airlineCode"`
	OperatingAirlineCode  string   `bson:"operatingAirlineCode" json:"operatingAirlineCode"`
	StopAirports          []string `bson:"stopAirports" json:"stopAirports"`
	DepartureTerminal     string   `bson:"departureTerminal" json:"departureTerminal"`
	ArrivalTerminal       string   `bson:"arrivalTerminal" json:"arrivalTerminal"`
}
