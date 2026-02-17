package smiles

import "time"

type Criteria struct {
	Adults                 string     `json:"adults"`
	CabinType              string     `json:"cabinType"`
	Children               string     `json:"children"`
	DepartureDate          time.Time  `json:"departureDate"`
	ReturnDate             *time.Time `json:"returnDate,omitempty"`
	DestinationAirportCode string     `json:"destinationAirportCode"`
	Infants                string     `json:"infants"`
	IsFlexibleDateChecked  string     `json:"isFlexibleDateChecked"`
	OriginAirportCode      string     `json:"originAirportCode"`
	TripType               string     `json:"tripType"`
	Region                 string     `json:"region"`
	CurrencyCode           string     `json:"currencyCode"`
	ForceCongener          string     `json:"forceCongener"`
	R                      string     `json:"r"`
}
