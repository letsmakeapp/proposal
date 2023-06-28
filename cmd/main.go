package main

import (
	"log"
	"proposal/internal/app/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const (
	Service     = "proposal"
	Environment = "development"
)

func initTraceProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(Service),
			attribute.String("environment", Environment),
		)),
	)
	return tp, nil
}

func initTracing(c *config.Variables) error {
	tp, err := initTraceProvider(c.Observation.JaegerEndpoint)
	if err != nil {
		return err
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return nil
}

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatalf("error while calling config.Read() = %s", err)
	}

	err = initTracing(&c)
	if err != nil {
		log.Fatalf("error while calling initTracing() = %s", err)
	}
}
