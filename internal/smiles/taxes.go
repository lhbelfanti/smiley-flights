package smiles

import (
	"context"
	"encoding/json"
	"net/url"

	"smiley-flights/internal/http"
	"smiley-flights/internal/log"
)

// GetTax retrieves boarding tax for a specific flight and fare.
type GetTax func(ctx context.Context, flight *Flight, fare *Fare, criteria Criteria) (*BoardingTax, error)

// MakeGetTax creates a new GetTax function
func MakeGetTax(httpClient http.Client, apiKey, domain string) GetTax {
	taxURL := "https://" + domain + "/v1/airlines/search"

	return func(ctx context.Context, flight *Flight, fare *Fare, criteria Criteria) (*BoardingTax, error) {
		u, _ := url.Parse(taxURL)
		q := u.Query()
		q.Set("adults", criteria.Adults)
		q.Set("children", criteria.Children)
		q.Set("infants", criteria.Infants)
		q.Set("uid", flight.UId)
		q.Set("fareuid", fare.UId)
		u.RawQuery = q.Encode()

		headers := map[string]string{
			"x-api-key":        apiKey,
			"region":           criteria.Region,
			"x-requested-with": "XMLHttpRequest",
			"authority":        domain,
			"user-agent":       "Mozilla/5.0",
		}

		resp, err := httpClient.NewRequest(ctx, "GET", u.String(), nil, headers)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteRequest
		}

		var data BoardingTax
		if err := json.Unmarshal([]byte(resp.Body), &data); err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToUnmarshalResponse
		}

		return &data, nil
	}
}
