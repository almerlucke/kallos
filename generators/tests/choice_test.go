package tests

import (
	"log"
	"testing"

	"github.com/almerlucke/kallos"
	"github.com/almerlucke/kallos/generators"
)

func TestChoice(t *testing.T) {
	c := generators.NewRandomChoice(kallos.ToValues(60, 61, 62, 63), false, false)

	index := 0
	for index < 20 {
		v := c.GenerateValue()
		log.Printf("v %v\n", v)
		index++
	}
}
