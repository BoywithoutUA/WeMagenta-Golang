package middlewares

import (
	"AIMusic-API/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func Tracer(serverName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cfg := &config.Configuration{
			ServiceName: serverName,
			Sampler: &config.SamplerConfig{
				Type:  "const",
				Param: 1,
			},
			Reporter: &config.ReporterConfig{
				LogSpans: true,
				LocalAgentHostPort: fmt.Sprintf("%s:%d",
					global.Config.Jaeger.Host, global.Config.Jaeger.Port),
			},
		}
		tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
		if err != nil {
			panic(fmt.Sprintf("Could not initialize jaeger tracer: %s", err.Error()))
		}
		defer closer.Close()

		span := tracer.StartSpan(ctx.Request.URL.Path)
		defer span.Finish()

		ctx.Set("tracer", tracer)
		ctx.Set("parentSpan", span)
		ctx.Next()
	}
}
