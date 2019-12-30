package tests

import (
	"testing"

	"github.com/almerlucke/kallos"
	"github.com/almerlucke/kallos/generators"
)

func TestRhythm(t *testing.T) {
	rhythm := kallos.NewRhythm(
		0.5,
		kallos.NewDurationStopCondition(10.0),
		generators.NewRandomChoice(kallos.ToValues(2.0, -2.0, 3.0, 1.0, -1.0), false, true),
	)

	sequence := rhythm.Generate()

	t.Logf("sequence %v\n", sequence)
}
