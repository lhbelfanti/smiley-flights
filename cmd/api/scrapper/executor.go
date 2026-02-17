package scrapper

import (
	"context"
	"fmt"
	"time"

	"smiley-flights/cmd/api/page"
	"smiley-flights/internal/log"
)

// Execute defines the function that starts the Smiles scraper
type Execute func(ctx context.Context, queryString string) error

// MakeExecute creates a new Execute
func MakeExecute(loadPage page.Load) Execute {
	return func(ctx context.Context, queryString string) error {
		log.Info(ctx, "Executing Smiles scrapper...")

		relativeURL := fmt.Sprintf("/emission?%s", queryString)

		err := loadPage(ctx, relativeURL, 60*time.Second)
		if err != nil {
			log.Error(ctx, fmt.Sprintf("failed to load Smiles search page: %v", err))
			return FailedToLoadSearchPage
		}

		// Placeholder for further scraping logic
		log.Info(ctx, "Smiles search page loaded successfully")

		return nil
	}
}
