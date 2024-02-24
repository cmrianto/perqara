package tracer

import (
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func NewTracer() (opentracing.Tracer, func(), error) {
	var t opentracing.Tracer
	var fnCloser func()
	tr := os.Getenv("TRACER")
	switch tr {
	case "datadog":
		dt := opentracer.New(
			tracer.WithServiceName(os.Getenv("SERVICE_NAME")),
			tracer.WithAnalytics(true),
			tracer.WithEnv(os.Getenv("ENV")),
			tracer.WithServiceVersion(os.Getenv("VERSION")),
		)
		fnCloser = func() {
			tracer.Stop()
		}
		t = dt
	case "jaeger":
		jt, closer, err := JaegerTracer(
			os.Getenv("SERVICE_NAME"),
			os.Getenv("DD_AGENT_HOST"),
			"",
		)
		if err != nil {
			log.Fatal("Cannot Init Jaeger Tracer", err.Error())
		}
		fnCloser = func() {
			closer.Close()
		}
		t = jt
	default:
		dt := mocktracer.New()
		fnCloser = func() {}
		t = dt
	}
	return t, fnCloser, nil
}
