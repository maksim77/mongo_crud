package main

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

type TracingContainer struct {
	*otlptrace.Exporter
	*sdktrace.TracerProvider
}

func (tc *TracingContainer) Stop(ctx context.Context) error {
	err := tc.ForceFlush(ctx)
	if err != nil {
		return fmt.Errorf("OTEL exporter flush failed: %w", err)
	}

	err = tc.Exporter.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("OTEL exporter shutdown failed: %w", err)
	}

	err = tc.TracerProvider.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("tracer provider shutdown failed: %w", err)
	}
	return nil
}

func initProvider() (*TracingContainer, error) {
	ctx := context.Background()
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("demo-client"),
		),
	)
	if err != nil {
		return nil, err
	}

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("0.0.0.0:4317"),
		otlptracegrpc.WithDialOption(grpc.WithBlock()))
	sctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	traceExp, err := otlptrace.New(sctx, traceClient)
	if err != nil {
		return nil, err
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExp)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tracerProvider)

	return &TracingContainer{
		Exporter:       traceExp,
		TracerProvider: tracerProvider,
	}, nil
}
