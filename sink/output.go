package sink

import "context"

// Producer defines the interface for sending data.
type Producer interface {
	Produce(ctx context.Context, data string) error
}
