package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"kafkaProjects/cmd/service_1/api"
	"kafkaProjects/cmd/service_1/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	connKafka, err := kafka.Dial("tcp", cfg.KafkaAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer connKafka.Close()

	topicConfigs := []kafka.TopicConfig{{
		Topic:             cfg.KafkaTopic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	},
	}
	err = connKafka.CreateTopics(topicConfigs...)
	if err != nil {
		log.Fatal(err)
	}

	kafkaWriter := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.KafkaAddr),
		Topic:                  cfg.KafkaTopic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	defer kafkaWriter.Close()
	h := api.NewHandler(kafkaWriter)

	router := http.NewServeMux()
	router.HandleFunc("POST /data", h.Data)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.PortService1), router)
	if err != nil {
		log.Fatal(err)
	}
}
