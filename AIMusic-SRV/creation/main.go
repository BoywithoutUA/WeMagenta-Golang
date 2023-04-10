package main

import (
	"creation/global"
	"creation/handler"
	"creation/proto"

	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 注册服务到Consul
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d",
		global.Config.Consul.Host, global.Config.Consul.Port)
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		panic(fmt.Sprintf("create consul client error: %v", err))
	}

	registration := &api.AgentServiceRegistration{
		ID:      "AIMusic-Creation-0",
		Name:    "AIMusic-Creation",
		Address: global.Config.SRVCreation.Host,
		Port:    global.Config.SRVCreation.Port,
		Tags:    []string{"Creation", "gRPC"},
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
	proto.RegisterCreationServer(server, &handler.CreationServer{})
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d",
		global.Config.SRVCreation.Host, global.Config.SRVCreation.Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(listener)
	if err != nil {
		panic("failed to start grpcSRV:" + err.Error())
	}
}
