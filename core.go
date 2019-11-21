package gokallos

// Value wraps a numeric type
type Value float64

// Generator generates a value if asked
type Generator interface {
	Value() Value
}
