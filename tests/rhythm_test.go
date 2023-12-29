package tests

import (
	"github.com/almerlucke/genny/bucket"
	"testing"

	"github.com/almerlucke/kallos"
)

func TestRhythm(t *testing.T) {
	rhythm := kallos.NewRhythm(
		0.5,
		kallos.NewDurationStopCondition(10.0),
		bucket.NewContinuous(bucket.Indexed, 2.0, -2.0, 3.0, 1.0, -1.0),
	)

	sequence := rhythm.Values()

	t.Logf("sequence %v\n", sequence)
}
