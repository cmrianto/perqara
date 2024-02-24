package tracer

import (
	"fmt"
	"io"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerUtils "github.com/uber/jaeger-client-go/utils"
)

func JaegerTracer(serviceName, agentHost, agentPort string) (opentracing.Tracer, io.Closer, error) {
	agentPortInt, _ := strconv.Atoi(agentPort)
	if agentPortInt == 0 {
		agentPortInt = jaeger.DefaultUDPSpanServerPort
	}
	jtransport, err := jaeger.NewUDPTransportWithParams(jaeger.UDPTransportParams{
		AgentClientUDPParams: jaegerUtils.AgentClientUDPParams{
			HostPort: fmt.Sprintf("%s:%d", agentHost, agentPortInt),
			Logger:   jaeger.StdLogger,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	reporter := jaeger.NewRemoteReporter(jtransport, jaeger.ReporterOptions.Logger(jaeger.StdLogger))
	reporter = jaeger.NewCompositeReporter(reporter, jaeger.NewLoggingReporter(jaeger.StdLogger))
	t, closer := jaeger.NewTracer(serviceName, jaeger.NewConstSampler(true), reporter, jaeger.TracerOptions.Logger(jaeger.StdLogger))
	return t, closer, nil
}
