package rhythm

import (
	"log"

	kallos "github.com/almerlucke/gokallos"

	"github.com/almerlucke/gokallos/generators/tools"
)

type Bouncer struct {
	durRamp   *tools.Ramp
	pauseRamp *tools.Ramp
	waitRamp  *tools.Ramp
	pause     bool
	reset     bool
}

func NewBouncer(durRamp *tools.Ramp, pauseRamp *tools.Ramp, waitRamp *tools.Ramp) *Bouncer {
	return &Bouncer{
		durRamp:   durRamp,
		pauseRamp: pauseRamp,
		waitRamp:  waitRamp,
	}
}

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

		log.Printf("wait %v\n", k)

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

func (b *Bouncer) IsContinuous() bool {
	return true
}

func (b *Bouncer) Done() bool {
	return false
}

func (b *Bouncer) Reset() {

}
