package kallos

// Stoppable can be stopped
type Stoppable interface {
	NumNoteEvents() int
	LastEventIsNote() bool
	Duration() float64
}

// StopCondition stops stream creation
type StopCondition interface {
	ShouldStop(s Stoppable) bool
}

// LengthStopCondition stops stream creation after the stream reaches a certain length
type LengthStopCondition struct {
	Length uint32
}

// NewLengthStopCondition returns a new length stop condition
func NewLengthStopCondition(length uint32) *LengthStopCondition {
	return &LengthStopCondition{
		Length: length,
	}
}

// ShouldStop return true if stream creation should stop
func (sc *LengthStopCondition) ShouldStop(s Stoppable) bool {
	return uint32(s.NumNoteEvents()) >= sc.Length
}

// DurationStopCondition stops stream creation after the stream reaches a certain duration in seconds
type DurationStopCondition struct {
	Duration float64
}

// NewDurationStopCondition returns a new duration stop condition
func NewDurationStopCondition(duration float64) *DurationStopCondition {
	return &DurationStopCondition{
		Duration: duration,
	}
}

// ShouldStop return true if stream creation should stop
func (sc *DurationStopCondition) ShouldStop(s Stoppable) bool {
	if s.LastEventIsNote() {
		return sc.Duration < (s.Duration() * SecondsPerBeat)
	}

	return false
}
