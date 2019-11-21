package generators

import (
	"math/rand"

	"github.com/almerlucke/gokallos"
)

// RandomChoice chooses a random element from the given values
type RandomChoice struct {
	Elements []gokallos.Value
}

// GenerateValue by choosing a random element
func (g *RandomChoice) GenerateValue() gokallos.Value {
	return g.Elements[rand.Intn(len(g.Elements))]
}
