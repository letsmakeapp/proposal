package webapi

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("webapi")
