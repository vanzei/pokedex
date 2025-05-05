package main

import (
	"context"
	"log"
	"time"

	"github.com/vanzei/pokedex/internal/pokeapi"
	"github.com/vanzei/pokedex/otel"
	"github.com/vanzei/pokedex/internal/pokecache"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Define the Prometheus counter globally
	apiCallCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_app_requests_total",
			Help: "Total number of requests to my app",
		},
		[]string{"method", "path"},
	)
)

func init() {
	// Register the Prometheus counter globally
	prometheus.MustRegister(apiCallCounter)
}

func main() {
	ctx := context.Background()

	// Initialize OpenTelemetry
	shutdown, err := otel.SetupOpenTelemetry(ctx)
	if err != nil {
		log.Fatalf("failed to set up OpenTelemetry: %v", err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Printf("failed to shut down OpenTelemetry: %v", err)
		}
	}()

	// Your application logic
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cache := pokecache.NewCache(5*time.Second) 

	cfg := &config{
		pokeapiClient:  pokeClient,
		apiCallCounter: apiCallCounter,
		cache:          cache,
		pokemonList:    make(map[string]pokeapi.PokemonResponse), // Initialize the map
	}

	startRepl(cfg)
}