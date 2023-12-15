package kallos

// Range to use in convert
type Range struct {
	Low  float64
	High float64
	Diff float64
}

// NewRange creates a new range
func NewRange(low float64, high float64) *Range {
	if high < low {
		tmp := high
		high = low
		low = tmp
	}

	return &Range{
		Low:  low,
		High: high,
		Diff: high - low,
	}
}

// ConvertShape convert a shape over a range
func (r *Range) ConvertShape(shape Shape, n int) Values {
	acc := 0.0
	inc := float64(len(shape)-1) / float64(n-1)
	result := Values{}

	dif := r.High - r.Low
	low := r.Low

	for n > 0 {
		result = append(result, Value{low + shape.Lookup(acc)*dif})
		acc += inc
		n--
	}

	return result
}
