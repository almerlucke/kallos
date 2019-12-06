package generators

import (
	kallos "github.com/almerlucke/gokallos"
)

// StaticValue generator
type StaticValue struct {
	Value kallos.Value
}

// NewStaticValue create a new static value
func NewStaticValue(v kallos.Value) *StaticValue {
	return &StaticValue{
		Value: v,
	}
}

// GenerateValue generates the value
func (v *StaticValue) GenerateValue() kallos.Value {
	return v.Value
}

// IsContinuous is always true
func (v *StaticValue) IsContinuous() bool {
	return true
}

// Done is always false
func (v *StaticValue) Done() bool {
	return false
}

// Reset does nothing
func (v *StaticValue) Reset() {
}
