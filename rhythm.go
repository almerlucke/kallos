package kallos

// Rhythm is here to be able to separate rhythm from a section if needed.
// We can precalculate rhythm with a stop condition and feed a sequential generator
// with the number of steps. This way we know the rhythm before creating a section,
// can be useful when we need to know the number of notes beforehand
type Rhythm struct {
	clock         float64
	until         StopCondition
	generator     Generator
	numNoteEvents int
	duration      float64
	rhythm        []float64
}

// NewRhythm initializes a new rhythm object
func NewRhythm(clock float64, until StopCondition, generator Generator) *Rhythm {
	return &Rhythm{
		clock:     clock,
		until:     until,
		generator: generator,
		rhythm:    []float64{},
	}
}

// Generate a rhythm until a condition is met
func (r *Rhythm) Generate() []float64 {
	for !r.until.ShouldStop(r) {
		// We have a fixed BPM of 120 (0.5 seconds per beat), calculate clock multiplier
		clockMultiplier := r.clock * BeatsPerSecond

		duration := r.generator.GenerateValue()[0] * clockMultiplier

		if duration < 0 {
			r.duration += -duration
		} else {
			r.duration += duration
			r.numNoteEvents++
		}

		r.rhythm = append(r.rhythm, duration)
	}

	return r.rhythm
}

// NumNoteEvents for stoppable interface
func (r *Rhythm) NumNoteEvents() int {
	return r.numNoteEvents
}

// LastEventIsNote for stoppable interface
func (r *Rhythm) LastEventIsNote() bool {
	if len(r.rhythm) > 1 {
		lastEvent := r.rhythm[len(r.rhythm)-1]
		return lastEvent >= 0.0
	}

	return false
}

// Duration for stoppable interface
func (r *Rhythm) Duration() float64 {
	return r.duration
}
