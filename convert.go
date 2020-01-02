package kallos

import "math"

// ShapeConverter can convert shapes into Values
type ShapeConverter interface {
	ConvertShape(shape Shape) Values
}

// Shape represents an abstract shape (between 0.0 and 1.0 inclusive)
// which can be used by a shape converter
type Shape []float64

// ConvertShape convert a shape to a slice of Value's,
// interpolate over the shape and do a lookup in Values
func (v Values) ConvertShape(shape Shape, n int) Values {
	acc := 0.0
	sl := len(shape)
	inc := float64(sl-1) / float64(n-1)
	m := float64(len(v) - 1)
	result := Values{}

	for n > 0 {
		si1 := int(acc)
		sf := acc - float64(si1)
		si2 := si1 + 1
		sv := shape[si1]

		if si2 < sl {
			sv += (shape[si2] - sv) * sf
		}

		result = append(result, v[int(math.Round(sv*m))])

		acc += inc
		n--
	}

	return result
}
