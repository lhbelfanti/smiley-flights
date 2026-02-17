package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"

	"smiley-flights/cmd/api/flights"
	_http "smiley-flights/internal/http"
	"smiley-flights/internal/log"
	"smiley-flights/internal/setup"
	"smiley-flights/internal/smiles"
)

const (
	smilesFlightsDomain = "api-air-flightsearch-green.smiles.com.ar"
	smilesTaxDomain     = "api-airlines-boarding-tax-green.smiles.com.ar"
)

func main() {
	/* --- Dependencies --- */
	ctx := context.Background()

	setup.Must(godotenv.Load())

	logLevel := zerolog.DebugLevel
	log.NewCustomLogger(os.Stdout, logLevel)

	httpClient := _http.NewClient()

	apiKey := os.Getenv("SMILES_API_KEY")
	authorization := os.Getenv("SMILES_AUTHORIZATION")

	// External Service
	getSmilesFlights := smiles.MakeGetFlights(httpClient, apiKey, smilesFlightsDomain, authorization)
	// getSmilesTax := smiles.MakeGetTax(httpClient, apiKey, smilesTaxDomain)

	// Services
	processResults := flights.MakeProcessResults() // getSmilesTax)
	searchFlights := flights.MakeSearch(getSmilesFlights, processResults)

	/* --- Router --- */
	log.Info(ctx, "Initializing router...")
	router := http.NewServeMux()

	router.HandleFunc("POST /flights/search/v1", flights.SearchHandlerV1(searchFlights))

	log.Info(ctx, "Router initialized!")

	/* --- Server --- */
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Info(ctx, fmt.Sprintf("smiley-flights server is ready to receive request on port %s", port))
	setup.Must(http.ListenAndServe(addr, router))
}
