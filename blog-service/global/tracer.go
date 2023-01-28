package global

import (
	"blog-service/pkg/tracer"
	"github.com/opentracing/opentracing-go"
)

var Tracer opentracing.Tracer

func SetupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service", "192.168.0.100:32687")
	if err != nil {
		return err
	}
	Tracer = jaegerTracer
	return nil
}