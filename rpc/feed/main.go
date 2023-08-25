package main

import (
	feed "douyin/kitex_gen/feed/feedservice"
	"log"
	"net"
	"github.com/cloudwego/kitex/server"
	
)


func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9991")
	svr := feed.NewServer(new(FeedServiceImpl),server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
