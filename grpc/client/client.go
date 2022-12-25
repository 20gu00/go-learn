package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-learn/pb"
	"log"
)

func main() {
	// 建立grpc连接
	// 改用WithTransportCredentials和insecure.NewCredentials（）
	a := insecure.NewCredentials()
	conn, err := grpc.Dial("127.0.0.1:8999", grpc.WithTransportCredentials(a)) // grpc.WithInsecure()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 关闭链接
	defer conn.Close()

	// 根据这个连接创建一个grpc的客户端
	client := pb.NewMessageSenderClient(conn)

	// 发起grpc请求,使用的server端的send方法
	resp, err := client.Send(context.Background(), &pb.MessageRequest{SaySomething: "hello world!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("receive message:", resp.GetResponseSomething())
}
