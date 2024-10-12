package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	PortService1 string `env:"PORT_SERVICE_1"`
	PortService2 string `env:"PORT_SERVICE_2"`
	KafkaAddr    string `env:"KAFKA_ADDR"`
	KafkaTopic   string `env:"KAFKA_TOPIC"`
}

func GetConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
