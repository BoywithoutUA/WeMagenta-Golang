package main

import (
	"administrator/global"
	"administrator/handler"
	"administrator/proto"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//// 创建配置对象
	//jaegerCfg := config.Configuration{
	//	ServiceName: "WeMagenta",
	//	Sampler: &config.SamplerConfig{
	//		Type:  "const",
	//		Param: 1,
	//	},
	//	Reporter: &config.ReporterConfig{
	//		LogSpans:            true,
	//		LocalAgentHostPort:  "127.0.0.1:6831", // Jaeger服务端的IP地址和端口
	//		BufferFlushInterval: 1,
	//	},
	//}
	//
	//// 创建Jaeger对象
	//jaegerMetricsFactory := metrics.NullFactory
	//jaegerTracer, _, err := jaegerCfg.NewTracer(
	//	config.Metrics(jaegerMetricsFactory),
	//)
	//if err != nil {
	//	log.Fatalf("failed to create Jaeger tracer: %v", err)
	//}
	//
	//// 设置全局跟踪器
	//opentracing.SetGlobalTracer(jaegerTracer)

	// 注册服务到Consul
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d",
		global.Config.Consul.Host, global.Config.Consul.Port)
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		panic(fmt.Sprintf("create consul client error: %v", err))
	}

	registration := &api.AgentServiceRegistration{
		ID:      "AIMusic-Administrator-0",
		Name:    "AIMusic-Administrator",
		Address: global.Config.SRVAdministrator.Host,
		Port:    global.Config.SRVAdministrator.Port,
		Tags:    []string{"Administrator", "gRPC"},
		Check: &api.AgentServiceCheck{
			HTTP: fmt.Sprintf("http://%s/consul/health",
				global.Config.Domain.Host),
			Timeout:  "5s",
			Interval: "30s",
			//DeregisterCriticalServiceAfter: "30s", // 如果检查失败，30秒后从Consul注销服务
		},
	}
	if err = consulClient.Agent().ServiceRegister(registration); err != nil {
		panic(fmt.Sprintf("register service error: %v", err))
	}

	server := grpc.NewServer()
	proto.RegisterAdministratorServer(server, &handler.AdministratorServer{})
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d",
		global.Config.SRVAdministrator.Host, global.Config.SRVAdministrator.Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(listener)
	if err != nil {
		panic("failed to start grpcSRV:" + err.Error())
	}
}
