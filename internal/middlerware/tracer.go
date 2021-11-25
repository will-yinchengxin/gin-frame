package middlerware

import (
	"context"
	"frame/consts"
	"frame/global"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Tracer() gin.HandlerFunc  {
	return func(c *gin.Context) {
		var ctx context.Context
		var span opentracing.Span

		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier{},
		)
		if err != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
			)

		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "Http"},
				)
		}

		// 日志信息追踪
		var tracerID string
		var spanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext := spanContext.(jaeger.SpanContext)
			tracerID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
		}

		c.Set(consts.XTraceID, tracerID)
		c.Set(consts.XSpanID, spanID)
		// c.Request = c.Request.WithContext(ctx) 使用这种方式会导致上下文丢失
		// 采用 c.Set 方式
		c.Set(consts.SpanFather, ctx)
		defer span.Finish()
		c.Next()
	}
}
