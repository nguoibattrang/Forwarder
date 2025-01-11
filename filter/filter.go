package filter

// Filter defines the interface for validating data.
type Filter interface {
	Validate(data string) bool
}
