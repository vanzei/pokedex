package main

import (
	"context"
	"log"
	"time"

	"github.com/vanzei/pokedex/internal/pokeapi"
    "github.com/vanzei/pokedex/otel"
)

func main() {
	ctx := context.Background()

	// Initialize OpenTelemetry
	shutdown, err := SetupOpenTelemetry(ctx)
	if err != nil {
		log.Fatalf("failed to set up OpenTelemetry: %v", err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Printf("failed to shut down OpenTelemetry: %v", err)
		}
	}()

	// Your application logic
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}