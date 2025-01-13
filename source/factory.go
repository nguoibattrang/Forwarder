package source

import (
	"fmt"
	"github.com/nguoibattrang/forwarder/config"
	"go.uber.org/zap"
)

// Create generates a Source instance based on the type and configuration.
func Create(sourceType string, cfg *config.ServiceConfig, log *zap.Logger) (Source, error) {
	switch sourceType {
	case "kafka":
		return NewKafkaSource(cfg.Source.Kafka, log), nil
	default:
		return nil, fmt.Errorf("unsupported source type: %s", sourceType)
	}
}
