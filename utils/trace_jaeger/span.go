package tracejaeger

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

type SpanCustomiser interface {
	customise() []trace.SpanStartOption
}

func NewSpan(ctx context.Context, serverName string, name string, cus SpanCustomiser) (context.Context, trace.Span) {
	if cus == nil {
		return otel.GetTracerProvider().Tracer(serverName).Start(ctx, name)
	}

	return otel.GetTracerProvider().Tracer(serverName).Start(ctx, name, cus.customise()...)
}

func GetContextWithMeta(ctx context.Context) (context.Context, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	traceIdString := md["x-trace-id"][0]
	traceId, err := trace.TraceIDFromHex(traceIdString)
	if err != nil {
		return nil, err
	}
	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceId,
	})
	ctx = trace.ContextWithSpanContext(ctx, spanContext)

	return ctx, nil
}

func SpanFromContext(ctx context.Context) trace.Span {
	return trace.SpanFromContext(ctx)
}

func AddSpanTags(span trace.Span, tags map[string]string) {
	list := make([]attribute.KeyValue, len(tags))

	var i int
	for k, v := range tags {
		list[i] = attribute.Key(k).String(v)
		i++
	}

	span.SetAttributes(list...)
}

func AddSpanEvents(span trace.Span, name string, events map[string]string) {
	list := make([]trace.EventOption, len(events))

	var i int
	for k, v := range events {
		list[i] = trace.WithAttributes(attribute.Key(k).String(v))
		i++
	}

	span.AddEvent(name, list...)
}

func AddSpanError(span trace.Span, err error) {
	span.RecordError(err)
}

func FailSpan(span trace.Span, msg string) {
	span.SetStatus(codes.Error, msg)
}
