package tests

import (
	"math"
	"testing"

	"github.com/almerlucke/kallos"
)

func TestConvert(t *testing.T) {
	shape := kallos.Shape{}

	i := 0.0

	for i < 20.0 {
		shape = append(shape, math.Sin(math.Pi*2*i/20.0)*0.5+0.5)
		i += 1.0
	}

	v := kallos.Values(kallos.ToValues(10, 9, 8, 7, 6, 5, 4, 3, 2, 1))

	cv := shape.Convert(v, 20)
	t.Logf("values %v\n", cv)

	r := kallos.NewRange(3, 30)

	cv = shape.Convert(r, 20).Apply(math.Round)
	t.Logf("range %v\n", cv)
}
