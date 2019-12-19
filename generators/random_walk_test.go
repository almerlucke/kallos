package generators

import (
	"log"
	"math/rand"
	"testing"
	"time"

	kallos "github.com/almerlucke/gokallos"
)

func TestRandomWalk(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	rand.Seed(seed)

	matrix := &RandomWalk2DMatrix{
		Values: [][]kallos.Value{
			kallos.ToValues(36, 38, 39, 42, 43, 44, 47),
			kallos.ToValues(48, 50, 51, 52, 54, 55, 58),
			kallos.ToValues(59, 60, 63, 64, 66, 65, 67),
			kallos.ToValues(69, 70, 73, 74, 75, 78, 79),
			kallos.ToValues(81, 82, 84, 86, 87, 90, 91),
			kallos.ToValues(92, 95, 96, 98, 99, 101, 103),
			kallos.ToValues(104, 107, 108, 109, 112, 113, 115),
		},
	}

	walker := NewRandomWalk([]int{7, 7}, matrix)

	cnt := 0
	for cnt < 20 {
		v := walker.GenerateValue()
		log.Printf("walker: %v\n", v)
		cnt++
	}
}
