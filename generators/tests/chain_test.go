package tests

import (
	"log"
	"testing"

	"github.com/almerlucke/kallos"
	"github.com/almerlucke/kallos/generators"
)

func TestChain(t *testing.T) {

	g1 := generators.NewSequence(kallos.ToValues(60, 61, 62, 63), false)
	g2 := generators.NewSequence(kallos.ToValues(80, 81, 82, 83), false)
	gc := generators.NewChain(g1, g2)
	gc.SetContinuous(false)

	index := 0
	for !gc.Done() && index < 20 {
		v := gc.GenerateValue()
		log.Printf("v %v\n", v)
		index++
	}
}
