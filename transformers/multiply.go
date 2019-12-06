package transformers

import (
	kallos "github.com/almerlucke/gokallos"
)

// Multiplier to multiply a section
type Multiplier struct {
	kallos.BasicTransformer
	Amount kallos.Value
}

// NewMultiplier creates a new multiplier
func NewMultiplier(v kallos.Value, t kallos.TransformType) *Multiplier {
	m := &Multiplier{
		Amount: v,
	}

	m.SetTransformType(t)

	return m
}

// TransformValue transform the value
func (m *Multiplier) TransformValue(v kallos.Value) kallos.Value {
	return v * m.Amount
}
