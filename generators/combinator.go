package generators

import (
	"github.com/almerlucke/kallos"
)

// Combinator combines generators in a single generator
type Combinator struct {
	Generators []kallos.Generator
}

// NewCombinator returns a new combinator
func NewCombinator(gens ...kallos.Generator) *Combinator {
	return &Combinator{
		Generators: gens,
	}
}

// GenerateValue by combining all values generated by the different generators,
// this can be used to form chords
func (c *Combinator) GenerateValue() kallos.Value {
	vs := kallos.Value{}

	for _, g := range c.Generators {
		tv := g.GenerateValue()
		vs = append(vs, tv...)
	}

	return vs
}

// IsContinuous for combinator returns false if any of the gens is not continuous
func (c *Combinator) IsContinuous() bool {
	isContinuous := true

	for _, g := range c.Generators {
		isContinuous = g.IsContinuous()
		if !isContinuous {
			break
		}
	}

	return isContinuous
}

// Done return true if any of the gens returns true
func (c *Combinator) Done() bool {
	done := false

	for _, g := range c.Generators {
		done = g.Done()
		if done {
			break
		}
	}

	return done
}

// Reset calls reset of gens
func (c *Combinator) Reset() {
	for _, g := range c.Generators {
		g.Reset()
	}
}
