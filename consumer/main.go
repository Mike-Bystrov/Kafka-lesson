package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

func main() {
	brockersStr := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_TOPIC")
	groupID := os.Getenv("KAFKA_GROUP_ID")
	consumerID := os.Getenv("CONSUMER_ID")

	brockers := strings.Split(brockersStr, ",")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brockers,
		GroupID:  groupID,
		Topic:    topic,
	})
	defer reader.Close()

	log.Printf("[%s] Consumer started in group [%s]", consumerID, groupID)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("[%s] Error reading message: %v", consumerID, err)
			break
		}
		log.Printf("[%s] Group: %s | Partition: %d | Offset: %d | Key: %s | Value: %s\n",
			consumerID, groupID, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
