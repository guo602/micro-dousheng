package main

import (
	feed "douyin/kitex_gen/feed/feedservice"
	"log"
	"net"
	"github.com/cloudwego/kitex/server"
	"douyin/pkg/config"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"	
	
)


func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.ServiceConfigInstance.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", config.ServiceConfigInstance.FeedService.Address)
	
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.FeedServiceName),
		provider.WithExportEndpoint(config.ExportEndpoint),
		provider.WithInsecure(),
	)
	svr := feed.NewServer(new(FeedServiceImpl),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServiceConfigInstance.FeedService.Name}), 
			server.WithServiceAddr(addr),
			server.WithSuite(tracing.NewServerSuite()),
			server.WithRegistry(r),server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
