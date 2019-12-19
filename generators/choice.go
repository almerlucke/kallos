package generators

import (
	"math/rand"

	kallos "github.com/almerlucke/gokallos"
)

// RandomChoice chooses a random element from the given values
type RandomChoice struct {
	elements []kallos.Value

	// In basket mode we get each next value from a shuffled copy of elements
	// so each value is picked once from the basket until the bucket is empty
	basketMode bool

	// needed if in basket mode
	basket []kallos.Value
	index  int
	loop   bool
	done   bool
}

// NewRandomChoice creates a new random choice generator
func NewRandomChoice(elements []kallos.Value, basketMode bool, loop bool) *RandomChoice {
	choice := &RandomChoice{
		elements:   elements,
		basketMode: basketMode,
		loop:       loop,
	}

	if basketMode {
		choice.basket = make([]kallos.Value, len(elements))
		copy(choice.basket, elements)
		rand.Shuffle(len(elements), func(i, j int) {
			choice.basket[i], choice.basket[j] = choice.basket[j], choice.basket[i]
		})
	}

	return choice
}

// GenerateValue by choosing a random element
func (g *RandomChoice) GenerateValue() kallos.Value {
	var v kallos.Value

	if g.basketMode {
		if g.done {
			v = g.basket[len(g.basket)-1]
		} else {
			v = g.basket[g.index]

			g.index++

			if g.index >= len(g.basket) {
				if g.loop {
					g.index = 0
					rand.Shuffle(len(g.basket), func(i, j int) {
						g.basket[i], g.basket[j] = g.basket[j], g.basket[i]
					})
				} else {
					g.done = true
				}
			}
		}
	} else {
		v = g.elements[rand.Intn(len(g.elements))]
	}

	return v
}

// IsContinuous for random choice is true if not in basket mode or else if loop is set
func (g *RandomChoice) IsContinuous() bool {
	return !g.basketMode || g.loop
}

// Done always returns done
func (g *RandomChoice) Done() bool {
	return g.done
}

// Reset index and reshuffle if basket mode
func (g *RandomChoice) Reset() {
	g.index = 0

	if g.basketMode {
		rand.Shuffle(len(g.basket), func(i, j int) {
			g.basket[i], g.basket[j] = g.basket[j], g.basket[i]
		})
	}
}
