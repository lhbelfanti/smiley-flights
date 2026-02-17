package main

import (
	"context"
	"fmt"
	"os"
	"smiley-flights/cmd/api/search"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"

	"smiley-flights/cmd/api/scrapper"
	"smiley-flights/internal/log"
	"smiley-flights/internal/setup"
	"smiley-flights/internal/webdriver"
)

func main() {
	ctx := context.Background()

	setup.Must(godotenv.Load())

	logLevel := zerolog.DebugLevel
	log.NewCustomLogger(os.Stdout, logLevel)

	log.Info(ctx, "Initializing WebDriver Manager...")
	newWebDriverManager := webdriver.MakeNewManager()
	wdManager := newWebDriverManager(ctx)
	defer setup.Must(wdManager.Quit(ctx))

	log.Info(ctx, "Initializing Smiles Scrapper...")
	makeScrapper := scrapper.MakeNew()
	executeScrapper := makeScrapper(wdManager.WebDriver())

	log.Info(ctx, "Executing Scrapper...")
	searchParameters := search.Parameters{
		Origin:         "BUE",
		Destination:    "CTG",
		DepartureDate:  "2026-07-04",
		ReturnDate:     "2026-07-15",
		Adults:         2,
		Children:       0,
		Infants:        0,
		IsFlexibleDate: false,
		TripType:       1,
		CabinType:      "all",
		CurrencyCode:   "ARS",
	}

	err := executeScrapper(ctx, searchParameters.ToQueryString())
	if err != nil {
		log.Error(ctx, fmt.Sprintf("Error executing scrapper: %v", err))
		return
	}

	time.Sleep(10 * time.Minute) // Wait time to visually understand what happened
}
