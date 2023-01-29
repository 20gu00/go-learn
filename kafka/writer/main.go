package main

import (
	"context"
	"time"
)

func main() {
	w := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{Address},
		Topic:    "t1",
		Parttion: 1,
		MinBytes: 1,
		MaxBytes: 1e6,
		MaxWait:  1000 * time.Millisecond,
	})
	//defer w.Close()
	err := w.WriterMessages(
		context.TODO(),
		kafka.Message{
			Parttion: 1,
			Key:      nil,
			Value:    []byte("hello"),
		},
		// kafka.Message{}
	)
}
