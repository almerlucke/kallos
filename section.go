package kallos

import (
	"github.com/almerlucke/genny"
	midi "github.com/almerlucke/gomidi"
)

// Section is a stream producer that uses generators for the production of stream events
type Section struct {
	Clock    float64
	Until    StopCondition
	Rhythm   genny.Generator[float64]
	Pitch    genny.Generator[Value]
	Velocity genny.Generator[Value]
	Channel  genny.Generator[int]
}

// Stream creates a new stream from the section
func (s *Section) Stream() *Stream {
	stream := NewStream()

	// We have a fixed BPM of 120 (0.5 seconds per beat), calculate clock multiplier
	clockMultiplier := s.Clock * BeatsPerSecond

	for !s.Until.ShouldStop(stream) {
		var event StreamEvent

		duration := s.Rhythm.NextValue() * clockMultiplier

		if duration < 0 {
			// Pause
			event = NewPause(-duration)
		} else if duration > 0 {
			// Note
			pitch := s.Pitch.NextValue()
			velocity := s.Velocity.NextValue()
			channel := s.Channel.NextValue()

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

// MidiTrack convert section to midi track
func (s *Section) MidiTrack() *midi.Track {
	return s.TimedNotes(0).MidiTrack()
}

// SequentialSection sequential sections
type SequentialSection []*Section

// MidiTrack to midi track
func (ss SequentialSection) MidiTrack() *midi.Track {
	stream := NewStream()

	for _, section := range ss {
		stream = stream.Append(section.Stream())
	}

	return stream.TimedNotes(0).MidiTrack()
}

// TimedSectionEntry holds start time and section
type TimedSectionEntry struct {
	// StartTime in seconds
	StartTime float64
	Section   *Section
}

// TimedSection play timed sections
type TimedSection []*TimedSectionEntry

// MidiTrack timed section to single midi track
func (ts TimedSection) MidiTrack() *midi.Track {
	notes := TimedNotes{}

	for _, entry := range ts {
		// Start time is in seconds, 120 BPM is default, 0.5 sec per beat,
		// so multiply with 2 to go from seconds to beats
		startTimeBeats := entry.StartTime * BeatsPerSecond
		notes = append(notes, entry.Section.TimedNotes(startTimeBeats)...)
	}

	return notes.MidiTrack()
}
