package kallos

import (
	"math"
)

const (
	// TicksPerQuarterNote constant
	TicksPerQuarterNote float64 = 192.0
	// BPM constant
	BPM float64 = 120.0
	// SecondsPerBeat constant
	SecondsPerBeat float64 = 0.5
	// BeatsPerSecond constant
	BeatsPerSecond float64 = 2.0
)

// Value wraps a generated value slice
type Value []float64

// Values wraps a slice of Value's
type Values []Value

// Shape represents an abstract shape (between 0.0 and 1.0 inclusive)
// which can be used by a shape converter
type Shape []float64

// ShapeConverter converts a shape to a slice of values
type ShapeConverter interface {
	ConvertShape(shape Shape, n int) Values
}

// IntegerRange to use in convert
type IntegerRange struct {
	Low  int64
	High int64
}

// Convert a shape
func (shape Shape) Convert(c ShapeConverter, n int) Values {
	return c.ConvertShape(shape, n)
}

// Lookup with a fractional index
func (shape Shape) Lookup(index float64) float64 {
	i1 := int(index)
	i2 := i1 + 1
	v := shape[i1]

	if i2 < len(shape) {
		v += (shape[i2] - v) * (index - float64(i1))
	}

	return v
}

// ConvertShape for a slice of Value's, interpolate over shape in n steps,
// use shape value to lookup index of Values
func (v Values) ConvertShape(shape Shape, n int) Values {
	acc := 0.0
	inc := float64(len(shape)-1) / float64(n-1)
	m := float64(len(v) - 1)
	result := Values{}

	for n > 0 {
		result = append(result, v[int(math.Round(shape.Lookup(acc)*m))])
		acc += inc
		n--
	}

	return result
}

// ConvertShape convert a shape over an integer range
func (r *IntegerRange) ConvertShape(shape Shape, n int) Values {
	acc := 0.0
	inc := float64(len(shape)-1) / float64(n-1)
	result := Values{}

	dif := float64(r.High - r.Low)
	min := float64(r.Low)

	for n > 0 {
		result = append(result, Value{math.Round(min + shape.Lookup(acc)*dif)})
		acc += inc
		n--
	}

	return result
}

// NewIntegerRange creates a new integer range
func NewIntegerRange(low int64, high int64) *IntegerRange {
	if high < low {
		tmp := high
		high = low
		low = tmp
	}

	return &IntegerRange{
		Low:  low,
		High: high,
	}
}

// Clip value to low and high
func Clip(v float64, low float64, high float64) float64 {
	if low > high {
		return math.Min(math.Max(float64(v), float64(high)), float64(low))
	}

	return math.Min(math.Max(float64(v), float64(low)), float64(high))
}

// ToValues convert a number of floats to a slice of Values
func ToValues(v ...float64) Values {
	vs := make(Values, len(v))

	for i, vc := range v {
		vs[i] = Value{vc}
	}

	return vs
}
