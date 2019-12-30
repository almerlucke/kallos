package generators

import (
	"log"
	"math/rand"
	"testing"

	"github.com/almerlucke/kallos"
)

func TestChoice(t *testing.T) {

	// seed := time.Now().UTC().UnixNano()

	rand.Seed(12232)

	c := NewRandomChoice(kallos.ToValues(60, 61, 62, 63), false, false)

	index := 0
	for index < 20 {
		v := c.GenerateValue()
		log.Printf("v %v\n", v)
		index++
	}
}
