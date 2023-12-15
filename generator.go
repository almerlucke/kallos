package kallos

// Generator generates a value if asked
// Done and Reset are ignored by most use cases
type Generator interface {
	// GenerateValue generate a value
	GenerateValue() Value
	// IsContinuous returns false if done and reset should be used,
	// this allows for generators that can have an end state
	IsContinuous() bool
	// Done checks if the generator is done generating values
	Done() bool
	// Reset the value generator
	Reset()
}
