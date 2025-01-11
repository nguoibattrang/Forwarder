package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// KafkaConfig holds configuration for Kafka source.
type KafkaConfig struct {
	g
	Address []string `yaml:"address"`
	Topic   string   `yaml:"topic"`
	Group   string   `yaml:"group"`
}

type SourceConfig struct {
	Type  string       `yaml:"type"`
	Kafka *KafkaConfig `yaml:"kafka"`
}

type SinkConfig struct {
	URL       string `yaml:"url"`
	SecretKey string `yaml:"secret_key"`
}

type Config struct {
	Source *SourceConfig `yaml:"source"`
	Sink   *SinkConfig   `yaml:"sink"`
}

// Parse parses the YAML configuration file into the Config struct.
func Parse(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
