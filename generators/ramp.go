package generators

import (
	"github.com/almerlucke/kallos/generators/tools"

	"github.com/almerlucke/kallos"
)

// Ramp generator wraps a ramp
type Ramp struct {
	ramp         *tools.Ramp
	isContinuous bool
	done         bool
}

// NewRamp initializes a new ramp
func NewRamp(ramp *tools.Ramp, isContinuous bool) *Ramp {
	return &Ramp{
		ramp:         ramp,
		isContinuous: isContinuous,
	}
}

// GenerateValue for a ramp generator
func (r *Ramp) GenerateValue() kallos.Value {
	v, d := r.ramp.Generate()
	if !d {
		if r.isContinuous {
			r.ramp.Reset()
		} else {
			r.done = true
		}
	}

	return kallos.Value{v}
}

// IsContinuous if continuous
func (r *Ramp) IsContinuous() bool {
	return r.isContinuous
}

// Done if done
func (r *Ramp) Done() bool {
	return r.done
}

// Reset if needed
func (r *Ramp) Reset() {
	r.done = false
	r.ramp.Reset()
}
