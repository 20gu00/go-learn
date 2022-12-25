package service

import (
	"context"
	"grpc-learn/pb"
	"log"
)

// 设置server端
type MessageSenderServer struct {
	*pb.UnimplementedMessageSenderServer // 未实施的
}

// 只
func (MessageSenderServer) Send(ctx context.Context, request *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Println("receive message:", request.GetSaySomething())
	resp := &pb.MessageResponse{}
	resp.ResponseSomething = "roger that!"
	return resp, nil
}
