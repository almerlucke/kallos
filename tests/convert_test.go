package tests

import (
	"github.com/almerlucke/kallos/plot"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/almerlucke/kallos"
)

func TestConvert(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	rand.Seed(seed)

	shape := kallos.Shape{}

	i := 0.0

	for i < 90.0 {
		shape = append(shape, math.Sin(math.Pi*2*i/90.0)*0.5+0.5)
		i += 1.0
	}

	v := kallos.ToValues(20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)

	cv := shape.Convert(v, 30)
	t.Logf("values %v\n", cv)

	plot.Plot(cv, "/Users/almerlucke/Desktop/shape.png", 500, 200)

	r := kallos.NewRange(3, 30)

	cv = shape.Convert(r, 20).Apply(math.Round)
	t.Logf("range %v\n", cv)

	s1 := kallos.Shape{0.0, 0.1, 0.4, 0.4, 0.4, 0.1, 0.0}
	s2 := kallos.Shape{0.01, 0.3, 0.5, 0.7, 0.5, 0.3, 0.01}
	m := kallos.NewMask(s1, s2)

	cv = m.Convert(100, kallos.NewRange(10, 50), func() float64 {
		return rand.Float64()
	})

	plot.Plot(cv, "/Users/almerlucke/Desktop/test.png", 500, 200)

	t.Logf("mask %v\n", cv)
}
