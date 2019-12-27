package generators

import (
	"github.com/almerlucke/gokallos/generators/tools"

	kallos "github.com/almerlucke/gokallos"
)

type Ramp struct {
	ramp         *tools.Ramp
	isContinuous bool
	done         bool
}

func NewRamp(ramp *tools.Ramp, isContinuous bool) *Ramp {
	return &Ramp{
		ramp:         ramp,
		isContinuous: isContinuous,
	}
}

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

func (r *Ramp) IsContinuous() bool {
	return r.isContinuous
}

func (r *Ramp) Done() bool {
	return r.done
}

func (r *Ramp) Reset() {
	r.done = false
	r.ramp.Reset()
}
