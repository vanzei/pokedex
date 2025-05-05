package otel

import (
    "context"
    "errors"
    "time"
	"net/http"
	"log"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/metric"
    "go.opentelemetry.io/otel/sdk/trace"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
    // Create a stdout metric exporter
    stdoutExporter, err := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
    if err != nil {
        return nil, err
    }

    // Create a channel to signal when the HTTP server is ready
    serverReady := make(chan struct{})

    // Expose Prometheus metrics via an HTTP endpoint
    go func() {
        mux := http.NewServeMux()
        mux.Handle("/metrics", promhttp.Handler()) // Use promhttp.Handler to expose metrics
        log.Println("Prometheus metrics available at http://localhost:8080/metrics")
        close(serverReady)
        log.Fatal(http.ListenAndServe(":8080", mux))
    }()

    // Wait for the server to be ready before proceeding
    <-serverReady

    // Create a MeterProvider for stdoutmetric
    meterProvider := metric.NewMeterProvider(
        metric.WithReader(metric.NewPeriodicReader(stdoutExporter, metric.WithInterval(100*time.Second))),
    )
    otel.SetMeterProvider(meterProvider)

    // Return a shutdown function
    return meterProvider.Shutdown, nil
}

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
        var combinedErr error
        for _, fn := range shutdownFuncs {
            if err := fn(ctx); err != nil {
                combinedErr = errors.Join(combinedErr, err)
            }
        }
        return combinedErr
    }

    return shutdown, nil
}