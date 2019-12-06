package generators

import (
	kallos "github.com/almerlucke/gokallos"
)

// Arpeggio play a scale of notes
type Arpeggio struct {
	StartNote   int
	Scale       []int
	Octaves     []int
	scaleIndex  int
	octaveIndex int
	loop        bool
	done        bool
}

// NewArpeggio creates a new arpeggio
func NewArpeggio(startNote int, scale []int, octaves []int, loop bool) *Arpeggio {
	return &Arpeggio{
		Scale:     scale,
		StartNote: startNote,
		Octaves:   octaves,
		loop:      loop,
	}
}

// GenerateValue generate a value
func (a *Arpeggio) GenerateValue() kallos.Value {
	if a.done {
		return kallos.Value(a.StartNote + a.Scale[len(a.Scale)-1] + a.Octaves[len(a.Octaves)-1]*12)
	}

	value := kallos.Value(a.StartNote + a.Scale[a.scaleIndex] + a.Octaves[a.octaveIndex]*12)

	a.scaleIndex++
	if a.scaleIndex >= len(a.Scale) {
		a.scaleIndex = 0
		a.octaveIndex++
		if a.octaveIndex >= len(a.Octaves) {
			if a.loop {
				a.octaveIndex = 0
			} else {
				a.done = true
			}
		}
	}

	return kallos.Value(value)
}

// IsContinuous is true if loop is set
func (a *Arpeggio) IsContinuous() bool {
	return a.loop
}

// Done return false if loop otherwise true when done
func (a *Arpeggio) Done() bool {
	return a.done
}

// Reset arpeggio
func (a *Arpeggio) Reset() {
	if !a.loop {
		a.scaleIndex = 0
		a.octaveIndex = 0
		a.done = false
	}
}
