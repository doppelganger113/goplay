package main

import (
	"context"
	"goplay.com/m/v2/examples/kafka_streaming"
	"log"
	"time"
)

func main() {
	topic := "golang-topic"

	err := kafka_streaming.StartProducer(
		context.Background(), "localhost:9092", topic, 10,
	)
	if err != nil {
		log.Println(err)
	}

	go func() {
		consumerErr := kafka_streaming.StartConsumer(context.Background(), topic)
		if consumerErr != nil {
			log.Println(consumerErr)
		}
	}()

	time.Sleep(15 * time.Second)
}
