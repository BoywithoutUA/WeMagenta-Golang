package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
	"user/global"
	"user/handler"
	"user/proto"
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
		ID:      "AIMusic-User-0",
		Name:    "AIMusic-User",
		Address: global.Config.SRVUser.Host,
		Port:    global.Config.SRVUser.Port,
		Tags:    []string{"User", "gRPC"},
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
	proto.RegisterUserServer(server, &handler.UserServer{})
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d",
		global.Config.SRVUser.Host, global.Config.SRVUser.Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(listener)
	if err != nil {
		panic("failed to start grpcSRV:" + err.Error())
	}
}
