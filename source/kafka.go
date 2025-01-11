package source

import (
	"context"
	"forwarder/config"
	"go.uber.org/zap"

	"forwarder/logger"
	"github.com/segmentio/kafka-go"
)

var log = logger.Log

// KafkaSource implements the Source interface for Kafka.
type KafkaSource struct {
	address     []string
	topic       string
	consumerGrp string
}

// NewKafkaSource creates a new KafkaSource.
func NewKafkaSource(config *config.KafkaConfig) *KafkaSource {

	return &KafkaSource{address: config.Address, topic: config.Topic, consumerGrp: config.Group}
}

// Consume starts consuming messages from Kafka.
func (k *KafkaSource) Consume(ctx context.Context) <-chan string {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: k.address,
		Topic:   k.topic,
		GroupID: k.consumerGrp,
	})

	out := make(chan string)

	go func() {
		defer r.Close()
		defer close(out)

		log.Info("Kafka consumer started", zap.String("topic", k.topic), zap.String("group", k.consumerGrp))
		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				log.Error("Error reading Kafka message", zap.Error(err))
				return
			}
			log.Debug("Kafka message received", zap.String("message", string(m.Value)))
			out <- string(m.Value)
		}
	}()

	return out
}
