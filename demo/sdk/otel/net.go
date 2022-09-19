package otel

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// Package-level tracer.
// This should be configured in your code setup instead of here.
var OtelNetApp = otel.Tracer("otel/net")

// sleepy mocks work that your application does.
func sleepy(ctx context.Context) {
	_, span := OtelNetApp.Start(ctx, "sleep")
	defer span.End()

	sleepTime := 1 * time.Second
	time.Sleep(sleepTime)
	span.SetAttributes(attribute.Int("sleep.duration", int(sleepTime)))
}

// httpHandler is an HTTP handler function that is going to be instrumented.
func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! I am instrumented automatically!")
	ctx := r.Context()
	sleepy(ctx)
}

func (o O) NetAppOfOtel() {
	// Wrap your httpHandler function.
	handler := http.HandlerFunc(httpHandler)
	http.Handle("/hello-instrumented", otelhttp.NewHandler(handler, "hello-instrumented"))

	// And start the HTTP serve.
	log.Fatal(http.ListenAndServe(":3030", nil))
}
