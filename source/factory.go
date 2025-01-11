package source

import (
	"fmt"
	"forwarder/config"
)

// Create generates a Source instance based on the type and configuration.
func Create(sourceType string, config *config.Config) (Source, error) {
	switch sourceType {
	case "kafka":
		return NewKafkaSource(config.Source.Kafka), nil
	default:
		return nil, fmt.Errorf("unsupported source type: %s", sourceType)
	}
}
