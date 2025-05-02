// filepath: /home/leovanzei/projects/bootdevprojects/pokedex/main.go
package otel

import (
    "context"
    "errors"
    "log"
    "net/http"
    "time"

    "github.com/vanzei/pokedex/internal/pokeapi"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/prometheus"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/metric"
    "go.opentelemetry.io/otel/sdk/trace"
)

var apiCallCounter = metric.Must(otel.GetMeter("pokeapi")).NewInt64Counter("api_calls")

// setupTracer initializes the OpenTelemetry tracer provider.
func setupTracer() (func(context.Context) error, error) {
    // Set up trace exporter
    traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
    if err != nil {
        return nil, err
    }

    // Create a TracerProvider
    tracerProvider := trace.NewTracerProvider(trace.WithBatcher(traceExporter))
    otel.SetTracerProvider(tracerProvider)

    // Return a shutdown function
    return tracerProvider.Shutdown, nil
}

// setupMeter initializes the OpenTelemetry meter provider.
func setupMeter() (func(context.Context) error, error) {
    // Create a Prometheus exporter
    prometheusExporter, err := prometheus.New()
    if err != nil {
        return nil, err
    }

    // Create a MeterProvider
    meterProvider := metric.NewMeterProvider(metric.WithReader(prometheusExporter))
    otel.SetMeterProvider(meterProvider)

    // Expose metrics via an HTTP endpoint
    go func() {
        http.Handle("/metrics", prometheusExporter)
        log.Println("Prometheus metrics available at http://localhost:8080/metrics")
        log.Fatal(http.ListenAndServe(":8080", nil))
    }()

    // Return a shutdown function
    return meterProvider.Shutdown, nil
}

// setupOpenTelemetry initializes both tracing and metrics.
func SetupOpenTelemetry(ctx context.Context) (func(context.Context) error, error) {
    var shutdownFuncs []func(context.Context) error

    // Initialize tracer
    tracerShutdown, err := setupTracer()
    if err != nil {
        return nil, err
    }
    shutdownFuncs = append(shutdownFuncs, tracerShutdown)

    // Initialize meter
    meterShutdown, err := setupMeter()
    if err != nil {
        return nil, err
    }
    shutdownFuncs = append(shutdownFuncs, meterShutdown)

    // Combined shutdown function
    shutdown := func(ctx context.Context) error {
        var err error
        for _, fn := range shutdownFuncs {
            err = errors.Join(err, fn(ctx))
        }
        return err
    }

    return shutdown, nil
}

func (c *Client) ListLocations(ctx context.Context, pageURL *string) (RespShallowLocations, error) {
    apiCallCounter.Add(ctx, 1)

    // Existing logic...
}

func main() {
    ctx := context.Background()

    // Initialize OpenTelemetry
    shutdown, err := setupOpenTelemetry(ctx)
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