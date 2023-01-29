package main

import (
	"context"
	"fmt"
)

// provide
const (
	Topic     = "t1"
	Address   = "127.0.0.1:9092"
	Partition = 1
)

func main() {
	conn, err := kafka.DialLeader(context.TODO(), "tcp", Address, Topic, Partition)
	if err != nil {
		//
		return
	}
	defer conn.Close()
	num, err := conn.WriteMessages(kafka.Message{Value: []byte("one")})
	fmt.Println(num)
	if err != nil {
		return
	}
	return
}
