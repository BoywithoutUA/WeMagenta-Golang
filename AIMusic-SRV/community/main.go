package main

import (
	"community/global"
	"community/handler"
	"community/proto"

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
		ID:      "AIMusic-Community-0",
		Name:    "AIMusic-Community",
		Address: global.Config.SRVCommunity.Host,
		Port:    global.Config.SRVCommunity.Port,
		Tags:    []string{"Community", "gRPC"},
		Check: &api.AgentServiceCheck{
			HTTP: fmt.Sprintf("http://%s/consul/health",
				global.Config.Domain.Host),
			Timeout:  "3s",
			Interval: "30s",
			//DeregisterCriticalServiceAfter: "30s", // 如果检查失败，30秒后从Consul注销服务
		},
	}
	if err = consulClient.Agent().ServiceRegister(registration); err != nil {
		panic(fmt.Sprintf("register service error: %v", err))
	}

	server := grpc.NewServer()
	proto.RegisterCommunityServer(server, &handler.CommunityServer{})
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d",
		global.Config.SRVCommunity.Host, global.Config.SRVCommunity.Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(listener)
	if err != nil {
		panic("failed to start grpcSRV:" + err.Error())
	}
}
