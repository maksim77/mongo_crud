package main

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func FollowSpan(ctx context.Context, name string) (_ context.Context, span trace.Span) {
	ctx, span = tracer.Start(ctx, name)
	return ctx, span
}
