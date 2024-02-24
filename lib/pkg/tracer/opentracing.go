package tracer

import (
	"context"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var defaultXRequestIDKey = "x-request-id"

type requestIDKey struct{}
type HttpRequestTracingFunc func(*http.Request, string) (*http.Request, *nethttp.Tracer)

func GrpcOpentracingUnaryClientInterceptor(t opentracing.Tracer, traceName string) grpc.UnaryClientInterceptor {
	return grpc_opentracing.UnaryClientInterceptor(
		grpc_opentracing.WithTracer(t),
		grpc_opentracing.WithTraceHeaderName(traceName),
	)
}

func GrpcOpentracingStreamClientInterceptor(t opentracing.Tracer, traceName string) grpc.StreamClientInterceptor {
	return grpc_opentracing.StreamClientInterceptor(
		grpc_opentracing.WithTracer(t),
		grpc_opentracing.WithTraceHeaderName(traceName),
	)
}

func GrpcOpentracingUnaryServerInterceptor(t opentracing.Tracer) grpc.UnaryServerInterceptor {
	return grpc_opentracing.UnaryServerInterceptor(
		grpc_opentracing.WithTracer(t),
	)
}

func GinOpentracingMiddleware(t opentracing.Tracer) gin.HandlerFunc {
	return ginhttp.Middleware(t)
}

func SetErrorSpan(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("event", "error", "message", err)
}

func HttpClientRequestWrapper(tr opentracing.Tracer) HttpRequestTracingFunc {
	return func(r *http.Request, traceName string) (*http.Request, *nethttp.Tracer) {
		req, ht := nethttp.TraceRequest(tr, r,
			nethttp.ComponentName(traceName),
			nethttp.ClientTrace(false),
		)
		return req, ht
	}
}

func RequestIdUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestID := handleRequestID(ctx)
		ctx = context.WithValue(ctx, requestIDKey{}, requestID)
		return handler(ctx, req)
	}
}

func handleRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return newRequestID()
	}

	header, ok := md[defaultXRequestIDKey]
	if !ok || len(header) == 0 {
		return newRequestID()
	}

	requestID := header[0]
	if requestID == "" {
		return newRequestID()
	}

	return requestID
}

func newRequestID() string {
	return uuid.New().String()
}

func RequestIdInjectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		if reqId := requestid.Get(c); reqId != "" {
			if sp := opentracing.SpanFromContext(ctx); sp != nil {
				sp.SetTag("request_id", reqId)
			}
			ctx = context.WithValue(ctx, requestIDKey{}, reqId)
			ctx = metadata.AppendToOutgoingContext(ctx, defaultXRequestIDKey, reqId)
		}
		c.Request = c.Request.WithContext(ctx)
	}
}
