package main

import (
	"net/rpc"
	"GoTest/src/imooc.com/ccmouse/learngo/lang/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error :%v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
