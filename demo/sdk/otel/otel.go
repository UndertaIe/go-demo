package otel

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	tracer "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	AppName   = "fib"
	TRACE_LOG = "./logs/traces.log"
)

// trace : span<Run> -> span<Poll> -> span<Write> -> span<Fibonacci>
// ====================================================================

//  ========================= span<Run> ===============================
//  ====== span<Poll> ======
//                          =============== span<Write> ===============
//                          ====== span<Fibonacci> ======

func (o O) MainOfOpenTelemety() {
	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create(TRACE_LOG)
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	// stdout exporter
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(f), // 说实话不支持函数命名参数有点尴尬，常用Option模式创建不定参数初始化的对象
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
	if err != nil {
		l.Fatal(err)
	}
	resource.Empty()
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	tp := tracer.NewTracerProvider(
		tracer.WithBatcher(exporter),
		tracer.WithResource(r),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	app := NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background()) //
	}()

	select {
	case <-sigCh: // 捕获中断信号
		tp.ForceFlush(context.Background())
		l.Println("\ngoodbye")
		return
	case err := <-errCh: // 接收chan信息
		tp.ForceFlush(context.Background()) // immediately exports all spans
		if err != nil {
			l.Fatal(err)
		}
	}
}

type App struct {
	r io.Reader
	l *log.Logger
}

// NewApp returns a new App.
func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r: r, l: l}
}

// Run starts polling users for Fibonacci number requests and writes results.
func (a *App) Run(ctx context.Context) error {
	for {
		newCtx, span := otel.Tracer(AppName).Start(ctx, "Run")

		n, err := a.Poll(newCtx)
		if err != nil {
			span.End()
			return err
		}
		a.Write(newCtx, n)
		span.End()
	}
}

// Poll asks a user for input and returns the request.
func (a *App) Poll(ctx context.Context) (uint, error) {
	_, span := otel.Tracer(AppName).Start(ctx, "Poll")
	defer span.End()

	a.l.Print("What Fibonacci number would you like to know: ")

	var n uint
	_, err := fmt.Fscanf(a.r, "%d\n", &n)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return 0, err
	}

	// Store n as a string to not overflow an int64.
	nStr := strconv.FormatUint(uint64(n), 10)
	span.SetAttributes(attribute.String("request.n", nStr))

	return n, nil
}

// Write writes the n-th Fibonacci number back to the user.
func (a *App) Write(ctx context.Context, n uint) {
	var span trace.Span
	ctx, span = otel.Tracer(AppName).Start(ctx, "Write")
	defer span.End()

	f, err := func(ctx context.Context) (uint64, error) { // 对工作代码无侵入性
		_, span := otel.Tracer(AppName).Start(ctx, "Fibonacci")
		defer span.End()
		f, err := Fibonacci(n)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		return f, err
	}(ctx)
	if err != nil {
		a.l.Printf("Fibonacci(%d): %v\n", n, err)
	} else {
		a.l.Printf("Fibonacci(%d) = %d\n", n, f)
	}
}

func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}
	if n > 93 {
		return 0, fmt.Errorf("unsupported fibonacci number %d: too large", n)
	}
	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}
