package main

import (
	"log"
	user "douyin/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"net"
	"github.com/cloudwego/kitex/server"
	"douyin/config"
	etcd "github.com/kitex-contrib/registry-etcd"

	

)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.ServiceConfigInstance.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", config.ServiceConfigInstance.UserService.Address)
	svr := user.NewServer(new(UserServiceImpl),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServiceConfigInstance.UserService.Name}), 
			server.WithServiceAddr(addr),
			server.WithRegistry(r),
		)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
