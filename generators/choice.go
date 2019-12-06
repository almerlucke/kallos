package generators

import (
	"math/rand"

	kallos "github.com/almerlucke/gokallos"
)

// RandomChoice chooses a random element from the given values
type RandomChoice struct {
	Elements []kallos.Value
}

// GenerateValue by choosing a random element
func (g *RandomChoice) GenerateValue() kallos.Value {
	return g.Elements[rand.Intn(len(g.Elements))]
}

// IsContinuous for random choice is true
func (g *RandomChoice) IsContinuous() bool {
	return true
}

// Done always returns false
func (g *RandomChoice) Done() bool {
	return false
}

// Reset does nothing
func (g *RandomChoice) Reset() {
}
