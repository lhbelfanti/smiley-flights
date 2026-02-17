package flights

import (
	"time"

	"smiley-flights/internal/smiles"
)

type (
	// FlightRequestDTO represents the search parameters for flights
	FlightRequestDTO struct {
		Origin              string `json:"origin"`
		Destination         string `json:"destination"`
		Departure           string `json:"departure"`
		Return              string `json:"return"`
		DaysBeforeDeparture int    `json:"daysBeforeDeparture"`
		DaysAfterDeparture  int    `json:"daysAfterDeparture"`
		DaysBeforeReturn    int    `json:"daysBeforeReturn"`
		DaysAfterReturn     int    `json:"daysAfterReturn"`
		Adults              string `json:"adults"`
	}

	// StopResponseDTO represents a flight stop
	StopResponseDTO struct {
		Airport string `json:"airport"`
		Hours   int    `json:"hours"`
		Minutes int    `json:"minutes"`
	}

	// FlightResponseDTO represents a single flight result
	FlightResponseDTO struct {
		Origin        string            `json:"origin"`
		Destination   string            `json:"destination"`
		Date          time.Time         `json:"date"`
		DepartureDate time.Time         `json:"departure_date"`
		ArrivalDate   time.Time         `json:"arrival_date"`
		FlightNumber  string            `json:"flight_number"`
		AirportFrom   string            `json:"airport_from"`
		AirportTo     string            `json:"airport_to"`
		Stops         int               `json:"stops"`
		AirportStops  []StopResponseDTO `json:"airport_stops,omitempty"`
		Cabin         string            `json:"cabin"`
		Airline       string            `json:"airline"`
		Baggage       int               `json:"baggage"`
		Miles         int               `json:"miles"`
		Tax           float32           `json:"tax"`
	}

	// SearchResponseDTO represents the full search response
	SearchResponseDTO struct {
		Departures []FlightResponseDTO `json:"departures"`
		Returns    []FlightResponseDTO `json:"returns"`
	}
)

func ToFlightResponseDTO(flight *smiles.Flight, fare *smiles.Fare, tax float32) FlightResponseDTO {
	var flightNumber string
	if len(flight.LegList) > 0 {
		flightNumber = flight.LegList[0].FlightNumber
	}

	var airportStops []StopResponseDTO
	if flight.Stops > 0 && len(flight.LegList) > 1 {
		// Stops info can be inferred from LegList
		// If there are N legs, there are N-1 stops.
		for i := 0; i < len(flight.LegList)-1; i++ {
			arrivalAtStop := flight.LegList[i].Arrival.Date
			departureFromStop := flight.LegList[i+1].Departure.Date
			duration := departureFromStop.Sub(arrivalAtStop)

			airportStops = append(airportStops, StopResponseDTO{
				Airport: flight.LegList[i].Arrival.Airport.Code,
				Hours:   int(duration.Hours()),
				Minutes: int(duration.Minutes()) % 60,
			})
		}
	}

	return FlightResponseDTO{
		Origin:        flight.Departure.Airport.Code,
		Destination:   flight.Arrival.Airport.Code,
		Date:          flight.Departure.Date,
		DepartureDate: flight.Departure.Date,
		ArrivalDate:   flight.Arrival.Date,
		FlightNumber:  flightNumber,
		AirportFrom:   flight.Departure.Airport.Code,
		AirportTo:     flight.Arrival.Airport.Code,
		Stops:         flight.Stops,
		AirportStops:  airportStops,
		Cabin:         flight.Cabin,
		Airline:       flight.Airline.Name,
		Baggage:       flight.Baggage.Quantity,
		Miles:         fare.Miles,
		Tax:           tax,
	}
}
