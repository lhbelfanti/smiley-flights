package flights

import (
	"time"

	"smiley-flights/internal/smiles"
)

type (
	// FlightRequestDTO represents the search parameters for flights
	FlightRequestDTO struct {
		Origin                string `json:"origin"`
		Destination           string `json:"destination"`
		Departure             string `json:"departure"`
		Return                string `json:"return"`
		DaysBeforeDeparture   int    `json:"daysBeforeDeparture"`
		DaysAfterDeparture    int    `json:"daysAfterDeparture"`
		DaysBeforeReturn      int    `json:"daysBeforeReturn"`
		DaysAfterReturn       int    `json:"daysAfterReturn"`
		Adults                string `json:"adults"`
		CabinType             string `json:"cabinType"`
		Children              string `json:"children"`
		Infants               string `json:"infants"`
		IsFlexibleDateChecked string `json:"isFlexibleDateChecked"`
		TripType              string `json:"tripType"`
		Region                string `json:"region"`
		CurrencyCode          string `json:"currencyCode"`
		ForceCongener         string `json:"forceCongener"`
		R                     string `json:"r"`
	}

	// FlightResponseDTO represents a single flight result
	FlightResponseDTO struct {
		Origin      string    `json:"origin"`
		Destination string    `json:"destination"`
		Date        time.Time `json:"date"`
		Cabin       string    `json:"cabin"`
		Airline     string    `json:"airline"`
		Stops       int       `json:"stops"`
		Miles       int       `json:"miles"`
		Tax         float32   `json:"tax"`
	}

	// SearchResponseDTO represents the full search response
	SearchResponseDTO struct {
		Departures []FlightResponseDTO `json:"departures"`
		Returns    []FlightResponseDTO `json:"returns"`
	}
)

func ToFlightResponseDTO(flight *smiles.Flight, fare *smiles.Fare, tax float32) FlightResponseDTO {
	return FlightResponseDTO{
		Origin:      flight.Departure.Airport.Code,
		Destination: flight.Arrival.Airport.Code,
		Date:        flight.Departure.Date,
		Cabin:       flight.Cabin,
		Airline:     flight.Airline.Name,
		Stops:       flight.Stops,
		Miles:       fare.Miles,
		Tax:         tax,
	}
}
