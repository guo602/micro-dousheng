package main

import (
	"log"
	user "douyin/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"net"
	"github.com/cloudwego/kitex/server"
	"douyin/pkg/config"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"	

	

)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.ServiceConfigInstance.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", config.ServiceConfigInstance.UserService.Address)

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.UserServiceName),
		provider.WithExportEndpoint(config.ExportEndpoint),
		provider.WithInsecure(),
	)

	svr := user.NewServer(new(UserServiceImpl),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServiceConfigInstance.UserService.Name}), 
			server.WithServiceAddr(addr),
			server.WithSuite(tracing.NewServerSuite()),
			server.WithRegistry(r),
		)

	
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
