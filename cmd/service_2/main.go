package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"kafkaProjects/cmd/service_1/config"
	"log"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.KafkaAddr},
		Topic:     cfg.KafkaTopic,
		Partition: 0,
		MaxBytes:  10e6,
	})

	defer kafkaReader.Close()

	for {
		err := func() error {
			ctx := context.Background()
			msg, err := kafkaReader.ReadMessage(ctx)
			if err != nil {
				return fmt.Errorf("read message:%w", err)
			}

			var sum int
			err = json.Unmarshal(msg.Value, &sum)
			if err != nil {
				return fmt.Errorf("unmarshal: %w", err)
			}
			fmt.Println(sum)
			return nil
		}()
		if err != nil {
			fmt.Errorf("%v", err)
			continue
		}
	}
}
