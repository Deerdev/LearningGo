package tracer

import (
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}

/*
上述代码主要分为三部分：

· config.Configuration：该项为 jaeger client 的配置项，主要设置应用的基本信息，如Sampler（固定采样、对所有数据都进行采样）、Reporter（是否启用LoggingReporter、刷新缓冲区的频率、上报的Agent地址）等。

· config.Configuration：该项为 jaeger client 的配置项，主要设置应用的基本信息，如Sampler（固定采样、对所有数据都进行采样）、Reporter（是否启用LoggingReporter、刷新缓冲区的频率、上报的Agent地址）等。

· cfg.NewTracer：根据配置项初始化Tracer对象，此处返回的是opentracing.Tracer，并不是某个供应商的追踪系统的对象。

· opentracing.SetGlobalTracer：设置全局的Tracer对象，根据实际情况设置即可。因为通常会统一使用一套追踪系统，因此该语句常常会被使用。
 */