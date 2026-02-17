package smiles

import (
	"context"
	"encoding/json"
	"net/url"

	"smiley-flights/internal/http"
	"smiley-flights/internal/log"
)

// GetFlights queries the Smiles API for flights.
type GetFlights func(ctx context.Context, criteria Criteria) (Data, error)

// MakeGetFlights creates a new GetFlights function
func MakeGetFlights(httpClient http.Client, apiKey, domain, authorization string) GetFlights {
	const dateLayout = "2006-01-02"
	flightsURL := "https://" + domain + "/v1/airlines/search"

	return func(ctx context.Context, criteria Criteria) (Data, error) {
		u, _ := url.Parse(flightsURL)
		q := u.Query()
		q.Set("adults", criteria.Adults)
		q.Set("cabinType", criteria.CabinType)
		q.Set("children", criteria.Children)
		q.Set("currencyCode", criteria.CurrencyCode)
		q.Set("departureDate", criteria.DepartureDate.Format(dateLayout))
		q.Set("destinationAirportCode", criteria.DestinationAirportCode)
		q.Set("infants", criteria.Infants)
		q.Set("isFlexibleDateChecked", criteria.IsFlexibleDateChecked)
		q.Set("originAirportCode", criteria.OriginAirportCode)
		if criteria.ReturnDate != nil {
			q.Set("returnDate", criteria.ReturnDate.Format(dateLayout))
		}
		q.Set("tripType", criteria.TripType)
		q.Set("forceCongener", criteria.ForceCongener)
		q.Set("r", criteria.R)
		u.RawQuery = q.Encode()

		headers := map[string]string{
			"accept":             "application/json, text/plain, */*",
			"accept-language":    "es-ES,es;q=0.9,en;q=0.8,de;q=0.7",
			"authorization":      authorization,
			"cache-control":      "no-cache",
			"channel":            "Web",
			"language":           "es-ES",
			"origin":             "https://www.smiles.com.ar",
			"pragma":             "no-cache",
			"priority":           "u=1, i",
			"referer":            "https://www.smiles.com.ar/",
			"region":             criteria.Region,
			"authority":          domain,
			"sec-ch-ua-platform": `"macOS"`,
			"sec-fetch-dest":     "empty",
			"sec-fetch-mode":     "cors",
			"sec-fetch-site":     "same-site",
			"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
			"x-api-key":          apiKey,
			"x-requested-with":   "XMLHttpRequest",
		}

		log.Info(ctx, u.String())

		resp, err := httpClient.NewRequest(ctx, "GET", u.String(), nil, headers)
		if err != nil {
			log.Error(ctx, err.Error())
			return Data{}, FailedToExecuteRequest
		}

		var data Data
		if err := json.Unmarshal([]byte(resp.Body), &data); err != nil {
			log.Error(ctx, err.Error())
			return Data{}, FailedToUnmarshalResponse
		}

		return data, nil
	}
}
