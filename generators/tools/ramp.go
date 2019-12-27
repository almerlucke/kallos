package tools

import "math"

type Ramp struct {
	min   float64
	dev   float64
	exp   float64
	steps int
	index int
	acc   float64
	inc   float64
}

func NewRamp(steps int, min float64, max float64, exp float64) *Ramp {
	return &Ramp{
		min:   min,
		dev:   max - min,
		steps: steps,
		exp:   exp,
		inc:   1.0 / float64(steps-1),
	}
}

func (r *Ramp) Reset() {
	r.acc = 0
	r.index = 0
}

func (r *Ramp) Generate() (float64, bool) {
	keepRunning := true

	v := math.Pow(r.acc, r.exp)*r.dev + r.min

	if r.index < (r.steps - 1) {
		r.acc += r.inc
		r.index++
	} else {
		keepRunning = false
	}

	return v, keepRunning
}
