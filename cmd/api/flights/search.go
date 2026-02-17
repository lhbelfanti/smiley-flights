package flights

import (
	"context"
	"sort"
	"time"

	"smiley-flights/internal/log"
	"smiley-flights/internal/smiles"
)

// Search queries the Smiles API for flights based on the request parameters.
type Search func(ctx context.Context, req FlightRequestDTO) (SearchResponseDTO, error)

// MakeSearch creates a new Search function
func MakeSearch(getFlights smiles.GetFlights, processResults ProcessResults) Search {
	const dateLayout = "2006-01-02"

	return func(ctx context.Context, req FlightRequestDTO) (SearchResponseDTO, error) {
		startingDepartureDate, err := time.Parse(dateLayout, req.Departure)
		if err != nil {
			log.Error(ctx, err.Error())
			return SearchResponseDTO{}, FailedToParseDepartureDate
		}

		startingReturningDate, err := time.Parse(dateLayout, req.Return)
		if err != nil {
			log.Error(ctx, err.Error())
			return SearchResponseDTO{}, FailedToParseReturnDate
		}

		criteria := smiles.Criteria{
			Adults:                 req.Adults,
			CabinType:              "all",
			Children:               "0",
			Infants:                "0",
			IsFlexibleDateChecked:  "false",
			TripType:               "2",
			Region:                 "ARGENTINA",
			CurrencyCode:           "ARS",
			ForceCongener:          "true",
			R:                      "ar",
			OriginAirportCode:      req.Origin,
			DestinationAirportCode: req.Destination,
		}

		var departureResults []smiles.Result
		var returnResults []smiles.Result

		for i := -req.DaysBeforeDeparture; i <= req.DaysAfterDeparture; i++ {
			departureDate := startingDepartureDate.AddDate(0, 0, i)

			c := criteria
			c.DepartureDate = departureDate
			c.ReturnDate = &startingReturningDate
			data, err := getFlights(ctx, c)
			if err != nil {
				continue
			}
			departureResults = append(departureResults, smiles.Result{Data: data, QueryDate: departureDate})
		}

		for i := -req.DaysBeforeReturn; i <= req.DaysAfterReturn; i++ {
			returnDate := startingReturningDate.AddDate(0, 0, i)

			c := criteria
			c.DepartureDate = returnDate
			c.ReturnDate = &startingDepartureDate
			c.OriginAirportCode = req.Destination
			c.DestinationAirportCode = req.Origin
			data, err := getFlights(ctx, c)
			if err != nil {
				continue
			}
			returnResults = append(returnResults, smiles.Result{Data: data, QueryDate: returnDate})
		}

		sortResults(departureResults)
		sortResults(returnResults)

		baseCriteria := smiles.Criteria{
			Adults:                req.Adults,
			Children:              "0",
			Infants:               "0",
			Region:                "ARGENTINA",
			CabinType:             "all",
			IsFlexibleDateChecked: "false",
			TripType:              "2",
			CurrencyCode:          "ARS",
			ForceCongener:         "true",
			R:                     "ar",
		}

		return SearchResponseDTO{
			Departures: processResults(ctx, departureResults, baseCriteria),
			Returns:    processResults(ctx, returnResults, baseCriteria),
		}, nil
	}
}

func sortResults(r []smiles.Result) {
	sort.Slice(r, func(i, j int) bool {
		return r[i].QueryDate.Before(r[j].QueryDate)
	})
}
