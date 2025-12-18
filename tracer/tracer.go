package tracer

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	otel_sdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func Initialize(ctx context.Context, config *Configuration) (trace.Tracer, error) {

	exporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(config.CollectorEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceName(config.ServiceName),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := otel_sdk.NewTracerProvider(
		otel_sdk.WithBatcher(exporter,
			otel_sdk.WithMaxExportBatchSize(config.ExportBatchSize),
			otel_sdk.WithBatchTimeout(time.Duration(config.BatchTimeOut)*time.Second),
			otel_sdk.WithMaxQueueSize(config.MaxQueueSize),
		),
		otel_sdk.WithResource(res),
		otel_sdk.WithSampler(otel_sdk.TraceIDRatioBased(config.SampleRatio)),
		otel_sdk.WithRawSpanLimits(otel_sdk.SpanLimits{
			AttributeCountLimit:         config.AttributeCountLimit,
			EventCountLimit:             config.EventCountLimit,
			LinkCountLimit:              config.LinkCountLimit,
			AttributePerEventCountLimit: config.AttributeCountLimit,
			AttributePerLinkCountLimit:  config.AttributePerLinkCountLimit,
		}),
	)

	otel.SetTracerProvider(tp)
	tracer := otel.Tracer(config.ServiceName)
	return tracer, nil
}
