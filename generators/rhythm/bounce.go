package rhythm

import (
	"github.com/almerlucke/kallos"

	"github.com/almerlucke/genny/floatgens/ramp"
)

// Bouncer represents a bouncing ball like rhythm
// duration ramp represents the hit the ground duration,
// pause ramp is the mid air duration
// wait ramp is the time between a new throw
type Bouncer struct {
	durRamp   *ramp.Ramp
	pauseRamp *ramp.Ramp
	waitRamp  *ramp.Ramp
	pause     bool
	reset     bool
}

// NewBouncer creates a new bouncer
func NewBouncer(durationRamp *ramp.Ramp, pauseRamp *ramp.Ramp, waitRamp *ramp.Ramp) *Bouncer {
	return &Bouncer{
		durRamp:   durationRamp,
		pauseRamp: pauseRamp,
		waitRamp:  waitRamp,
	}
}

// GenerateValue generate a value
func (b *Bouncer) NextValue() kallos.Value {
	var k float64

	if b.reset {
		b.pauseRamp.Reset()
		b.durRamp.Reset()
		b.pause = false
		b.reset = false

		k = b.waitRamp.NextValue()
		if b.waitRamp.Done() {
			b.waitRamp.Reset()
		}

		k = -k
	} else {
		if b.pause {
			k = b.pauseRamp.NextValue()
			if !b.pauseRamp.Done() {
				b.pause = !b.pause
			} else {
				b.reset = true
			}

			k = -k
		} else {
			k = b.durRamp.NextValue()
			if !b.durRamp.Done() {
				b.pause = !b.pause
			} else {
				b.reset = true
			}
		}
	}

	return kallos.Value{k}
}

// Continuous is true
func (b *Bouncer) Continuous() bool {
	return true
}

// Done is false
func (b *Bouncer) Done() bool {
	return false
}

// Reset is empty
func (b *Bouncer) Reset() {
}
