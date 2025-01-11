package main

import (
	"context"
	"forwarder/logger"
	"forwarder/sink"
	"forwarder/source"
	"go.uber.org/zap"

	"forwarder/config"
	"forwarder/transform"
)

func main() {
	//
	logger.InitLogger()
	log := logger.Log

	// Load configuration
	configFile := "E:\\Kaicode\\gomongo\\app.yml"
	cfg, err := config.Parse(configFile)
	if err != nil {
		log.Error("config.Parse failed to load config", zap.String("topic", configFile), zap.Error(err))
		return
	}

	// Create Input using Factory
	s, err := source.Create(cfg.Source.Type, cfg)
	if err != nil {
		log.Error("source.Create failed to create input source", zap.String("topic", cfg.Source.Type), zap.Error(err))
		return
	}
	messages := s.Consume(context.Background())

	// Initialize Transformer and Output
	transformer := transform.NewHTMLToMarkdown()
	producer := sink.NewDifyProducer(cfg.Sink)

	// Process Messages
	log.Info("Processing messages...")
	for message := range messages {
		transformedMessage := transformer.Transform(message)

		if err := producer.Produce(context.Background(), transformedMessage); err != nil {
			log.Error("producer.Produce Failed to send message", zap.Error(err))
		} else {
			log.Debug("Successfully sent message")
		}
	}
}
