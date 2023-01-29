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
	// 读取开始的偏移量
	first, err := conn.ReadFirstOffset()
	fmt.Println(first)
	if err != nil {
		return
	}
	// 读取最新的偏移量
	last, err := conn.ReadLastOffset()
	fmt.Println(last)
	if err != nil {
		return
	}
	// seek设置偏移量等待新消息
	newOffset, err := conn.Seek(last, kafka.SeekAbsolute)
	fmt.Println(newOffset)
	if err != nil {
		return
	}
	// TODO 消费者10分钟没有收到消息 ReadMessage会返回EOF error,可以稳定复现,reader
	for message, err := conn.ReadMessage(1e6)
	fmt.Println(message)
	if err != nil {
		return
	}
	)
	return
}
