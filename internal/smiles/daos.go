package smiles

import (
	"encoding/json"
	"time"
)

type Fare struct {
	UId   string `json:"uid"`
	FType string `json:"type"`
	Miles int    `json:"miles"`
	Money int    `json:"money"`
}

type Airline struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type FlightDetail struct {
	Date    time.Time `json:"date"`
	Airport Airport   `json:"airport"`
}

type Leg struct {
	Cabin        string       `json:"cabin"`
	Departure    FlightDetail `json:"departure"`
	Arrival      FlightDetail `json:"arrival"`
	FlightNumber string       `json:"flightNumber"`
	Duration     int          `json:"duration"`
}

type Baggage struct {
	Free     string `json:"free"`
	Quantity int    `json:"quantity"`
}

type Duration struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

type Flight struct {
	UId            string       `json:"uid"`
	Cabin          string       `json:"cabin"`
	Stops          int          `json:"stops"`
	Departure      FlightDetail `json:"departure"`
	Arrival        FlightDetail `json:"arrival"`
	Airline        Airline      `json:"airline"`
	Baggage        Baggage      `json:"baggage"`
	Duration       Duration     `json:"duration"`
	DurationNumber int          `json:"durationNumber"`
	TimeStop       Duration     `json:"timeStop"`
	LegList        []Leg        `json:"legList"`
	FareList       []Fare       `json:"fareList"`
}

type BestPricing struct {
	Miles      int    `json:"miles"`
	SourceFare string `json:"sourceFare"`
	Fare       Fare   `json:"fare"`
}

type Segment struct {
	SegmentType string      `json:"type"`
	FlightList  []Flight    `json:"flightList"`
	BestPricing BestPricing `json:"bestPricing"`
	Airports    Airports    `json:"airports"`
}

type Airport struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Airports struct {
	DepartureAirports []Airport `json:"departureAirportList"`
	ArrivalAirports   []Airport `json:"arrivalAirportList"`
}
type Data struct {
	RequestedFlightSegmentList []Segment `json:"requestedFlightSegmentList"`
}

type Result struct {
	Data      Data
	QueryDate time.Time
}

func (f *FlightDetail) UnmarshalJSON(p []byte) error {
	var aux struct {
		Date    string  `json:"date"`
		Airport Airport `json:"airport"`
	}

	err := json.Unmarshal(p, &aux)
	if err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02T15:04:05", aux.Date)
	if err != nil {
		return err
	}

	f.Date = t
	f.Airport = aux.Airport

	return nil
}
