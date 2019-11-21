package gokallos

// Generator generates a value if asked
type Generator interface {
	GenerateValue() Value
}
