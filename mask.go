package kallos

// Mask represents an abstract field which can be used to convert
type Mask struct {
	Low  Shape
	High Shape
	dif  float64
	min  float64
	max  float64
}

// NewMask creates a mask from two shapes
func NewMask(l Shape, h Shape) *Mask {
	m := &Mask{
		Low:  l,
		High: h,
	}

	low := 0.0
	high := 0.0

	for i, lv := range l {
		hv := h[i]

		if lv < low {
			low = lv
		}

		if hv < low {
			low = lv
		}

		if lv > high {
			high = lv
		}

		if hv > high {
			high = hv
		}
	}

	m.min = low
	m.max = high
	m.dif = high - low

	return m
}

// Convert mask to Values
func (m *Mask) Convert(n int, r *Range, dist func() float64) Values {
	acc := 0.0
	inc := float64(len(m.Low)-1) / float64(n-1)
	result := Values{}

	for n > 0 {
		l := (m.Low.Lookup(acc) - m.min) / m.dif
		h := (m.High.Lookup(acc) - m.min) / m.dif

		if l > h {
			t := h
			h = l
			l = t
		}

		v := ((h-l)*dist()+l)*r.Diff + r.Low

		result = append(result, Value{v})
		acc += inc
		n--
	}

	return result
}
