package main

import (
	"google.golang.org/grpc"
	"grpc-learn/pb"
	"grpc-learn/service"
	"log"
	"net"
)

// 根据实际的service服务开启grpc的server端
func main() {
	// 注册一个grpc服务
	srv := grpc.NewServer()
	// 注册实际的服务
	pb.RegisterMessageSenderServer(srv, service.MessageSenderServer{})
	// 开启一个tcp,rpc基于tcp,grpc基于http2,实际上还是tcp底层
	listener, err := net.Listen("tcp", ":8999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 开启grpc服务,提供服务端的服务  使用这个tcp listener提供服务
	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
