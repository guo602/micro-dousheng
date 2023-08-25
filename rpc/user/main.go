package main

import (
	"log"
	user "douyin/kitex_gen/user/userservice"
	"net"
	"github.com/cloudwego/kitex/server"
	

)

func main() {

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9990")
	svr := user.NewServer(new(UserServiceImpl),server.WithServiceAddr(addr))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
