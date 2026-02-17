package flights

import (
	"encoding/json"
	"net/http"

	"smiley-flights/internal/http/response"
	"smiley-flights/internal/log"
)

const (
	InvalidRequestBody = "Invalid request body"
	FailedToSearch     = "Failed to search flights"
)

// SearchHandlerV1 HTTP Handler of the endpoint /flights/search/v1
func SearchHandlerV1(search Search) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req FlightRequestDTO
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}

		ctx = log.With(ctx,
			log.Param("origin", req.Origin),
			log.Param("destination", req.Destination),
			log.Param("departure", req.Departure),
			log.Param("return", req.Return),
			log.Param("daysBeforeDeparture", req.DaysBeforeDeparture),
			log.Param("daysAfterDeparture", req.DaysAfterDeparture),
			log.Param("daysBeforeReturn", req.DaysBeforeReturn),
			log.Param("daysAfterReturn", req.DaysAfterReturn),
			log.Param("adults", req.Adults),
			log.Param("cabinType", req.CabinType),
			log.Param("children", req.Children),
			log.Param("infants", req.Infants),
			log.Param("isFlexibleDateChecked", req.IsFlexibleDateChecked),
			log.Param("tripType", req.TripType),
			log.Param("region", req.Region),
			log.Param("currencyCode", req.CurrencyCode),
			log.Param("forceCongener", req.ForceCongener),
			log.Param("r", req.R),
		)

		results, err := search(ctx, req)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToSearch, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Search completed successfully", results, nil)
	}
}
