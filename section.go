package kallos

import (
	"sort"

	midi "github.com/almerlucke/gomidi"
)

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
	return uint32(s.NumNoteEvents) >= sc.Length
}

// DurationStopCondition stops stream creation after the stream reaches a certain duration
type DurationStopCondition struct {
	Duration float64
}

// ShouldStop return true if stream creation should stop
func (sc *DurationStopCondition) ShouldStop(s *Stream) bool {
	if len(s.Events) > 1 {
		lastEvent := s.Events[len(s.Events)-1]
		if lastEvent.Type() == NoteEvent {
			return s.Duration >= sc.Duration
		}
	}

	return false
}

// Section is a stream producer that uses generators for the production of stream events
type Section struct {
	Clock    float64
	Until    StopCondition
	Rhythm   Generator
	Pitch    Generator
	Velocity Generator
	Channel  Generator
}

// Stream creates a new stream from the section
func (s *Section) Stream() *Stream {
	stream := NewStream()

	// We have a fixed BPM of 120 (0.5 seconds per beat), calculate clock multiplier
	clockMultiplier := s.Clock / 0.5

	for !s.Until.ShouldStop(stream) {
		var event StreamEvent

		duration := s.Rhythm.GenerateValue()[0] * clockMultiplier

		if duration < 0 {
			// Pause
			event = NewPause(-duration)
		} else if duration > 0 {
			// Note
			pitch := s.Pitch.GenerateValue()
			velocity := s.Velocity.GenerateValue()
			channel := int(s.Channel.GenerateValue()[0])

			event = NewNote(duration, pitch, velocity, channel)
		}

		stream.AddEvent(event)
	}

	stream.Sanitize()

	return stream
}

// TimedNotes get timed notes
func (s *Section) TimedNotes(startTime float64) TimedNotes {
	stream := s.Stream()

	return stream.TimedNotes(startTime)
}

// ToMidiTrack convert section to midi track
func (s *Section) ToMidiTrack(ticksPerQuarterNote float64) *midi.Track {
	return s.TimedNotes(0).ToMidiTrack(ticksPerQuarterNote)
}

// SequentialSection sequential sections
type SequentialSection []*Section

// ToMidiTrack to midi track
func (ss SequentialSection) ToMidiTrack(ticksPerQuarterNote float64) *midi.Track {
	stream := NewStream()

	for _, section := range ss {
		otherStream := section.Stream()
		stream = stream.Append(otherStream)
	}

	notes := stream.TimedNotes(0)
	sort.Sort(notes)

	notes.CalculateDeltaTimes()

	return notes.ToMidiTrack(ticksPerQuarterNote)
}

// TimedSectionEntry holds start time and section
type TimedSectionEntry struct {
	// StartTime in seconds
	StartTime float64
	Section   *Section
}

// TimedSection play timed sections
type TimedSection []*TimedSectionEntry

// ToMidiTrack timed section to single midi track
func (ts TimedSection) ToMidiTrack(ticksPerQuarterNote float64) *midi.Track {
	notes := TimedNotes{}

	for _, entry := range ts {
		// Start time is in seconds, 120 BPM is default, 0.5 sec per beat,
		// so multiply with 2 to go from seconds to beats
		startTimeBeats := entry.StartTime * 2.0
		notes = append(notes, entry.Section.TimedNotes(startTimeBeats)...)
	}

	sort.Sort(notes)

	notes.CalculateDeltaTimes()

	return notes.ToMidiTrack(ticksPerQuarterNote)
}
