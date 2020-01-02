package rhythm

import (
	"github.com/almerlucke/kallos"

	"github.com/almerlucke/kallos/generators/tools"
)

// Bouncer represents a bouncing ball like rhythm
// duration ramp represents the hit the ground duration,
// pause ramp is the mid air duration
// wait ramp is the time between a new throw
type Bouncer struct {
	durRamp   *tools.Ramp
	pauseRamp *tools.Ramp
	waitRamp  *tools.Ramp
	pause     bool
	reset     bool
}

// NewBouncer creates a new bouncer
func NewBouncer(durationRamp *tools.Ramp, pauseRamp *tools.Ramp, waitRamp *tools.Ramp) *Bouncer {
	return &Bouncer{
		durRamp:   durationRamp,
		pauseRamp: pauseRamp,
		waitRamp:  waitRamp,
	}
}

// GenerateValue generate a value
func (b *Bouncer) GenerateValue() kallos.Value {
	var k float64
	var d bool

	if b.reset {
		b.pauseRamp.Reset()
		b.durRamp.Reset()
		b.pause = false
		b.reset = false

		k, d = b.waitRamp.Generate()
		if !d {
			b.waitRamp.Reset()
		}

		k = -k
	} else {
		if b.pause {
			k, d = b.pauseRamp.Generate()
			if d {
				b.pause = !b.pause
			} else {
				b.reset = true
			}

			k = -k
		} else {
			k, d = b.durRamp.Generate()
			if d {
				b.pause = !b.pause
			} else {
				b.reset = true
			}
		}
	}

	return kallos.Value{k}
}

// IsContinuous is true
func (b *Bouncer) IsContinuous() bool {
	return true
}

// Done is false
func (b *Bouncer) Done() bool {
	return false
}

// Reset is empty
func (b *Bouncer) Reset() {
}
