package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serverName, hostPort string) (opentracing.Tracer, io.Closer, error) {
	defer func() {
		if err := recover(); err != nil {
			// todo
		}
	}()

	//if hostPort == "" {
	//	return errors.New("host should not empty")
	//}
	//
	//if ServiceName == "" {
	//	return errors.New("ServiceName should not empty")
	//}

	// 初始化配置项
	JaegerConfig := &jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{ // 固定采样, 对所有数据采样
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		ServiceName: serverName,
		Reporter: &jaegercfg.ReporterConfig{ // 是否启用 loggerReport, 缓冲区的频率, 以及上报的地址
			LogSpans:            true,
			LocalAgentHostPort:  hostPort,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	tracer, closer, err := JaegerConfig.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer) // 设置全局的 trace 对象, 可以根据实际情况设置
	return tracer, closer, nil
}
