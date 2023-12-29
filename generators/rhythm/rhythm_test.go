package rhythm

import (
	"github.com/almerlucke/genny/floatgens/ramp"
	"log"
	"testing"
)

func TestRhythm(t *testing.T) {

	r := NewBouncer(
		ramp.New(10, 0.075, 0.25, 1.4),
		ramp.New(10, 0.075, 0.25, 1.4),
		ramp.New(4, 2.0, 0.25, 1.4),
	)

	index := 0

	for index < 50 {
		v := r.NextValue()
		log.Printf("value %v\n", v)
		index++
	}
}
