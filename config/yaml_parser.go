package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// KafkaConfig holds configuration for Kafka source.
type KafkaConfig struct {
	Address []string `yaml:"address"`
	Topic   string   `yaml:"topic"`
	Group   string   `yaml:"group"`
}

type SourceConfig struct {
	Type  string       `yaml:"type"`
	Kafka *KafkaConfig `yaml:"kafka"`
}

type LogConfig struct {
	Mode string `yaml:"mode"`
}

type SinkConfig struct {
	URL       string `yaml:"url"`
	SecretKey string `yaml:"secret_key"`
}

type ServiceConfig struct {
	Source *SourceConfig `yaml:"source"`
	Sink   *SinkConfig   `yaml:"sink"`
	Logger *LogConfig    `yaml:"logger"`
}

func LoadConfig(filename string) (*ServiceConfig, error) {
	// Initialize viper
	v := viper.New()
	v.SetConfigFile(filename)
	v.SetConfigType("yaml")

	// Set defaults if needed
	v.SetDefault("logger.mode", "production")

	// Read in the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal into the ServiceConfig struct
	var config ServiceConfig
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
