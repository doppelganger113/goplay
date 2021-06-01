package kafka_streaming

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func StartProducer(ctx context.Context, address string, topic string, nMessages uint) error {
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		return err
	}
	defer func() {
		if connErr := conn.Close(); connErr != nil {
			log.Println(connErr)
		}
	}()

	var i uint
	for ; i < nMessages; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			_, err = conn.WriteMessages(
				kafka.Message{Value: []byte("My message")},
			)
			if err != nil {
				return err
			}
			log.Println("wrote message")
			time.Sleep(time.Second)
		}
	}

	return nil
}

func StartConsumer(ctx context.Context, topic string) error {
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		return err
	}
	defer func() {
		if connErr := conn.Close(); connErr != nil {
			log.Println(connErr)
		}
	}()

	if err = conn.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return err
	}

	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	b := make([]byte, 10e3)            // 10KB max per message
	for {
		_, err = batch.Read(b)
		if err != nil {
			log.Println("Failed reading", err)
			break
		}
		fmt.Println(string(b))
	}

	if err = batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	return nil
}
