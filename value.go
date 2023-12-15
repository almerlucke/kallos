package kallos

import (
	"gonum.org/v1/plot/plotter"
	"math"
)

// Value wraps a generated value slice
type Value []float64

// Values wraps a slice of Value's
type Values []Value

// ToFloat converts a slice of Value's back to a float64 slice
func (v Values) ToFloat() []float64 {
	fls := make([]float64, len(v))

	for i, e := range v {
		fls[i] = e[0]
	}

	return fls
}

// Apply a function to all elements (and sub elements) of Values
func (v Values) Apply(f func(float64) float64) Values {
	for _, e := range v {
		for i := range e {
			e[i] = f(e[i])
		}
	}

	return v
}

func (v Values) ApplyN(f func(float64, ...float64) float64, argv ...float64) Values {
	for _, e := range v {
		for i := range e {
			e[i] = f(e[i], argv...)
		}
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

// ToXYs returns a plotter XY's slice
func (v Values) ToXYs() plotter.XYs {
	pts := make(plotter.XYs, len(v))

	for i, v := range v {
		pts[i] = plotter.XY{
			X: float64(i + 1),
			Y: v[0],
		}
	}

	return pts
}

// ToValues convert a number of floats to a slice of Values
func ToValues(v ...float64) Values {
	vs := make(Values, len(v))

	for i, vc := range v {
		vs[i] = Value{vc}
	}

	return vs
}

// IntToValues convert a number of integers to a slice of Values
func IntToValues(v ...int) Values {
	vs := make(Values, len(v))

	for i, vc := range v {
		vs[i] = Value{float64(vc)}
	}

	return vs
}
