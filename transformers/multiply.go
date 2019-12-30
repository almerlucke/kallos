package transformers

import (
	"github.com/almerlucke/kallos"
)

// Multiplier to multiply a section
type Multiplier struct {
	kallos.BasicTransformer
	Amount float64
}

// NewMultiplier creates a new multiplier
func NewMultiplier(v float64, t kallos.TransformType) *Multiplier {
	m := &Multiplier{
		Amount: v,
	}

	m.SetTransformType(t)

	return m
}

// TransformValue transform the value
func (m *Multiplier) TransformValue(v kallos.Value) kallos.Value {
	vc := make(kallos.Value, len(v))

	for i, sv := range v {
		vc[i] = sv * m.Amount
	}

	return vc
}
