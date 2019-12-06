package gokallos

// StopCondition stops stream creation
type StopCondition interface {
	ShouldStop(s *Stream) bool
}

// LengthStopCondition stops stream creation after the stream reaches a certain length
type LengthStopCondition struct {
	Length uint32
}

// ShouldStop return true if stream creation should stop
func (sc *LengthStopCondition) ShouldStop(s *Stream) bool {
	return uint32(len(s.Events)) >= sc.Length
}

// DurationStopCondition stops stream creation after the stream reaches a certain duration
type DurationStopCondition struct {
	Duration Value
}

// ShouldStop return true if stream creation should stop
func (sc *DurationStopCondition) ShouldStop(s *Stream) bool {
	return s.Duration >= sc.Duration
}

// Section is a stream producer that uses generators for the production of stream events
type Section struct {
	Clock    Value
	Until    StopCondition
	Rhythm   Generator
	Pitch    Generator
	Velocity Generator
	Channel  Generator
}

// Stream creates a new stream from the section
func (s *Section) Stream() *Stream {
	stream := NewStream()

	for !s.Until.ShouldStop(stream) {
		var event StreamEvent

		duration := s.Rhythm.GenerateValue() * s.Clock

		if duration < 0 {
			// Pause
			event = NewPause(-duration)
		} else if duration > 0 {
			// Note
			pitch := s.Pitch.GenerateValue()
			velocity := s.Velocity.GenerateValue()
			channel := s.Channel.GenerateValue()

			event = NewNote(duration, pitch, velocity, channel)
		}

		stream.AddEvent(event)
	}

	stream.Sanitize()

	return stream
}
