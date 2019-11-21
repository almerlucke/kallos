package gokallos

import (
	"math"
)

// Value wraps a numeric type
type Value float64

// Clip value to low and high
func Clip(v Value, low Value, high Value) Value {
	return Value(math.Min(math.Max(float64(v), float64(low)), float64(high)))
}
