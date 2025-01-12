package source

import "context"

// Source defines the interface for consuming data from various sources.
type Source interface {
	Consume(ctx context.Context) <-chan Data
}
