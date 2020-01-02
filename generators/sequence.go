package generators

import (
	"github.com/almerlucke/kallos"
)

// Sequence of values
type Sequence struct {
	Values kallos.Values
	index  int
	loop   bool
	done   bool
}

// NewSequence returns a new sequence
func NewSequence(values kallos.Values, loop bool) *Sequence {
	return &Sequence{
		Values: values,
		index:  0,
		loop:   loop,
	}
}

// GenerateValue returns next value in sequence
func (s *Sequence) GenerateValue() kallos.Value {
	if s.done {
		return s.Values[len(s.Values)-1]
	}

	r := s.Values[s.index]

	s.index++

	if s.index >= len(s.Values) {
		if s.loop {
			s.index = 0
		} else {
			s.done = true
		}
	}

	return r
}

// IsContinuous returns true if not manual
func (s *Sequence) IsContinuous() bool {
	return s.loop
}

// Done can be asked by for instance Markov state to see if the generator is done generating for
// that state. To be able to use it with Sequence we added a manual indicator
func (s *Sequence) Done() bool {
	return s.done
}

// Reset the sequence, does nothing if mode is not manual
func (s *Sequence) Reset() {
	if !s.loop {
		s.index = 0
		s.done = false
	}
}
