package generators

import (
	kallos "github.com/almerlucke/gokallos"
)

// Sequence of values
type Sequence struct {
	Index        int
	Values       []kallos.Value
	isContinuous bool
}

// NewSequence returns a new sequence
func NewSequence(values []kallos.Value, isContinuous bool) *Sequence {
	return &Sequence{
		Values:       values,
		Index:        0,
		isContinuous: isContinuous,
	}
}

// GenerateValue returns next value in sequence
func (s *Sequence) GenerateValue() kallos.Value {
	r := s.Values[s.Index]

	if s.Index < (len(s.Values) - 1) {
		s.Index++
	} else if s.isContinuous {
		s.Index = 0
	}

	return r
}

// IsContinuous returns true if not manual
func (s *Sequence) IsContinuous() bool {
	return s.isContinuous
}

// Done can be asked by for instance Markov state to see if the generator is done generating for
// that state. To be able to use it with Sequence we added a manual indicator
func (s *Sequence) Done() bool {
	if s.isContinuous {
		return false
	}

	return s.Index >= (len(s.Values) - 1)
}

// Reset the sequence, does nothing if mode is not manual
func (s *Sequence) Reset() {
	if !s.isContinuous {
		s.Index = 0
	}
}
