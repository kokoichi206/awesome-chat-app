package util

import (
	"errors"
	"io"
	"net"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(host, port, service string) (opentracing.Tracer, io.Closer, error) {
	if host == "" || port == "" {
		return nil, nil, errors.New("host or port is empty.")
	}

	addr := net.JoinHostPort(host, port)

	cfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: time.Second,
			LocalAgentHostPort:  addr,
		},
	}

	//nolint:wrapcheck
	return cfg.NewTracer()
}
