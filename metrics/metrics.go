package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/zapr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.uber.org/zap"
)

func Initialize(ctx context.Context, config *Configuration) (*metric.MeterProvider, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(config.ServiceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	zapLogger, _ := zap.NewDevelopment()
	otel.SetLogger(zapr.NewLogger(zapLogger))

	promExporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create prometheus exporter: %w", err)
	}

	otlpExporter, err := otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpoint(config.CollectorEndpoint),
		otlpmetrichttp.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create otlp metric exporter: %w", err)
	}

	mp := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(promExporter),
		metric.WithReader(metric.NewPeriodicReader(otlpExporter,
			metric.WithInterval(time.Duration(config.ExportInterval)*time.Second),
		)),
	)

	otel.SetMeterProvider(mp)

	return mp, nil
}
