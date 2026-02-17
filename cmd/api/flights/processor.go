package flights

import (
	"context"

	"smiley-flights/internal/smiles"
)

// ProcessResults processes the search results to find the cheapest flight for each day.
type ProcessResults func(ctx context.Context, results []smiles.Result, criteria smiles.Criteria) []FlightResponseDTO

// MakeProcessResults creates a new ProcessResults function.
func MakeProcessResults() ProcessResults { // getTax smiles.GetTax) ProcessResults {
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
				/*
					boardingTax, err := getTax(ctx, cheapestFlightDay, cheapestFareDay, criteria)
					if err != nil {
						log.Warn(ctx, err.Error())
						// We don't want to abort processing due to a tax retrieval error
					}
				*/
				tax := float32(0)
				/*
					if err == nil && boardingTax != nil {
						tax = boardingTax.Totals.Total.Money
					}
				*/
				processedResults = append(processedResults, ToFlightResponseDTO(cheapestFlightDay, cheapestFareDay, tax))
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
