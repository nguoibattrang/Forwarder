package source

import (
	"context"
	"encoding/json"
	"github.com/nguoibattrang/forwarder/config"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

// KafkaSource implements the Source interface for Kafka.
type KafkaSource struct {
	address     []string
	topic       string
	consumerGrp string
	log         *zap.Logger
}

// NewKafkaSource creates a new KafkaSource.
func NewKafkaSource(config *config.KafkaConfig, log *zap.Logger) *KafkaSource {
	return &KafkaSource{address: config.Address, topic: config.Topic, consumerGrp: config.Group, log: log}
}

// Consume starts consuming messages from Kafka.
func (inst *KafkaSource) Consume(ctx context.Context) <-chan Data {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: inst.address,
		Topic:   inst.topic,
		GroupID: inst.consumerGrp,
	})

	out := make(chan Data)

	go func() {
		defer r.Close()
		defer close(out)

		inst.log.Info("Kafka consumer started", zap.String("topic", inst.topic), zap.String("group", inst.consumerGrp))
		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				inst.log.Error("Error reading Kafka message", zap.Error(err))
				return
			}
			var data Data
			err = json.Unmarshal(m.Value, &data)
			if err != nil {
				inst.log.Error("Error unmarshalling Kafka message", zap.Error(err))
				continue
			}
			out <- data
		}
	}()

	return out
}
