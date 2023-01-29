package main

import (
	"context"
	"log"
	"time"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{Address},
		Topic:    "t1",
		Parttion: 1,
		MinBytes: 1,
		MaxBytes: 1e6,
		MaxWait:  1000 * time.Millisecond,
	})
	defer r.Close()
	for (
		message, err := r.ReadMessage(context.TODO())
	// TODO可以设置重新连接
	if err != nil {
		return
	}
	log.Println(message.Topic, message.Parttion, message.Offset, string(message.Value))
	)

}