package sink

// Producer defines the interface for sending data.
type Producer interface {
	Produce(data string) error
}
