package rhythm

import (
	"log"
	"testing"

	"github.com/almerlucke/gokallos/generators/tools"
)

func TestRhythm(t *testing.T) {

	r := NewBouncer(
		tools.NewRamp(10, 0.075, 0.25, 1.4),
		tools.NewRamp(10, 0.075, 0.25, 1.4),
		tools.NewRamp(4, 2.0, 0.25, 1.4),
	)

	index := 0

	for index < 50 {
		v := r.GenerateValue()
		log.Printf("value %v\n", v)
		index++
	}
}
