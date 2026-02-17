package flights

import (
	"context"

	"smiley-flights/internal/smiles"
)

// ProcessResults processes the search results to find the cheapest flight for each day.
type ProcessResults func(ctx context.Context, results []smiles.Result, criteria smiles.Criteria) []FlightResponseDTO

// MakeProcessResults creates a new ProcessResults function.
func MakeProcessResults() ProcessResults {
	const bigMaxMilesNumber = 9_999_999

	return func(ctx context.Context, results []smiles.Result, criteria smiles.Criteria) []FlightResponseDTO {
		var processedResults []FlightResponseDTO

		for _, v := range results {
			if len(v.Data.RequestedFlightSegmentList) == 0 || len(v.Data.RequestedFlightSegmentList[0].FlightList) == 0 {
				continue
			}

			var cheapestFlightDay *smiles.Flight
			cheapestFareDay := &smiles.Fare{
				Miles: bigMaxMilesNumber,
			}

			for _, f := range v.Data.RequestedFlightSegmentList[0].FlightList {
				smilesClubFare := getSmilesClubFare(&f)
				if smilesClubFare != nil && cheapestFareDay.Miles > smilesClubFare.Miles {
					cheapestFlightDay = &f
					cheapestFareDay = smilesClubFare
				}
			}

			if cheapestFareDay.Miles != bigMaxMilesNumber {
				processedResults = append(processedResults, ToFlightResponseDTO(cheapestFlightDay, cheapestFareDay, 0))
			}
		}

		return processedResults
	}
}

func getSmilesClubFare(f *smiles.Flight) *smiles.Fare {
	for _, fare := range f.FareList {
		if fare.FType == "SMILES_CLUB" {
			return &fare
		}
	}
	return nil
}
