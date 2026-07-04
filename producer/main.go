package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	brockersStr := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_TOPIC")

	brockers := strings.Split(brockersStr, ",")

	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brockers...),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		Async:                  false,
	}
	defer writer.Close()

	log.Printf("Producer started! Topic: %s", topic)

	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("key: %d", t.Unix())),
			Value: []byte(fmt.Sprintf("Meassage at %s", t.Format(time.RFC3339))),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Printf("Error^ %v", err)
			continue
		}
		log.Println("Message sent successfully")
	}
}
