package otel

type O struct{}

func (o O) Name() string {
	return "otel"
}

func (o O) Desc() string {
	return "Understand how a system is functioning when it is failing or having issues is critical to resolving those issues. One strategy to understand this is with tracing. This guide shows how the OpenTelemetry Go project can be used to trace an example application. You will start with an application that computes Fibonacci numbers for users, and from there you will add instrumentation to produce tracing telemetry with OpenTelemetry Go."
}

func (o O) DemoOfJaeger() {

}



