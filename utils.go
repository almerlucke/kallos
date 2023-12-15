package kallos

import (
	"math"
)

// Clip value to low and high
func Clip(v float64, low float64, high float64) float64 {
	if low > high {
		return math.Min(math.Max(v, high), low)
	}

	return math.Min(math.Max(v, low), high)
}

// Round to a quantization unit
func Round(v float64, quantization float64) float64 {
	return math.Round(v/quantization) * quantization
}
