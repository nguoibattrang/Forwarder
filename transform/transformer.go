package transform

// Transformer defines the interface for transforming data.
type Transformer interface {
	Transform(data string) string
}
