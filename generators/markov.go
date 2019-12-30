package generators

import (
	"math/rand"

	"github.com/almerlucke/kallos"
)

// MarkovState for value and next state
type MarkovState struct {
	Value         kallos.Generator
	Probabilities []float64
	States        []*MarkovState
}

// NewMarkovState create a new markov state
func NewMarkovState(value kallos.Generator, probabilities []float64, states []*MarkovState) *MarkovState {
	return &MarkovState{
		Value:         value,
		Probabilities: probabilities,
		States:        states,
	}
}

// NextState from this state
func (s *MarkovState) NextState() *MarkovState {
	if len(s.States) > 0 {
		r := rand.Float64()
		pa := 0.0
		index := 0
		for _, p := range s.Probabilities {
			pa += p
			if r < pa {
				return s.States[index]
			}
			index++
		}
	}

	return nil
}

// MarkovChain is a collection of state transitions,
// each state can represent any value generator (allows for nested chains)
type MarkovChain struct {
	States       []*MarkovState
	Start        *MarkovState
	Current      *MarkovState
	isContinuous bool
}

// NewMarkovChain creates a new markov chain
func NewMarkovChain(states []*MarkovState, start *MarkovState, isContinuous bool) *MarkovChain {
	return &MarkovChain{
		States:       states,
		Start:        start,
		Current:      start,
		isContinuous: isContinuous,
	}
}

// GenerateValue value
func (c *MarkovChain) GenerateValue() kallos.Value {
	v := c.Current.Value.GenerateValue()

	if c.Current.Value.IsContinuous() {
		c.Current = c.Current.NextState()
	} else if c.Current.Value.Done() {
		c.Current.Value.Reset()
		c.Current = c.Current.NextState()
	}

	return v
}

// IsContinuous should return true if the chain does not have an end state
func (c *MarkovChain) IsContinuous() bool {
	return c.isContinuous
}

// Done to check if the chain is at it's end
func (c *MarkovChain) Done() bool {
	if c.isContinuous {
		return false
	}

	return c.Current == nil
}

// Reset the chain
func (c *MarkovChain) Reset() {
	if !c.isContinuous {
		c.Current = c.Start
	}
}
