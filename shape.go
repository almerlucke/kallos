package kallos

// Shape represents an abstract shape (between 0.0 and 1.0 inclusive)
// which can be used by a shape converter
type Shape []float64

// ShapeConverter converts a shape to a slice of values
type ShapeConverter interface {
	ConvertShape(shape Shape, n int) Values
}

// CreateShape by executing n times f(i, n)
func CreateShape(f func(int, int) float64, n int) Shape {
	i := 0
	shape := Shape{}

	for i < n {
		shape = append(shape, f(i, n))
		i++
	}

	return shape
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
