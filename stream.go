package gokallos

import (
	"fmt"
	"math"

	midi "github.com/almerlucke/gomidi"
)

// StreamEventType identifies the type of stream event
type StreamEventType int

const (
	// NoteEvent plays a note (or chord)
	NoteEvent StreamEventType = iota
	// PauseEvent take a pause
	PauseEvent
)

// StreamEvent is the building block of a stream
type StreamEvent interface {
	fmt.Stringer
	Duration() float64
	Type() StreamEventType
	ApplyTransform(transformer Transformer) StreamEvent
}

// basicEvent implements the full StreamEvent interface and can be used by other stream events
type basicEvent struct {
	duration  float64
	eventType StreamEventType
}

func (e *basicEvent) Duration() float64 {
	return e.duration
}

func (e *basicEvent) Type() StreamEventType {
	return e.eventType
}

// Note stream event
type Note struct {
	basicEvent
	Pitch    Value
	Velocity Value
	Channel  int
}

// NewNote creates an initialized note
func NewNote(duration float64, pitch Value, velocity Value, channel int) *Note {
	return &Note{
		basicEvent: basicEvent{
			duration:  duration,
			eventType: NoteEvent,
		},
		Pitch:    pitch,
		Velocity: velocity,
		Channel:  channel,
	}
}

// String representation
func (e *Note) String() string {
	return fmt.Sprintf("{duration: %v, pitch: %v, velocity: %v, channel: %v}", e.duration, e.Pitch, e.Velocity, e.Channel)
}

// ApplyTransform to note
func (e *Note) ApplyTransform(transformer Transformer) StreamEvent {
	ce := NewNote(e.duration, e.Pitch, e.Velocity, e.Channel)

	switch transformer.TransformType() {
	case TransformDuration:
		ce.duration = transformer.TransformValue(Value{ce.duration})[0]
	case TransformPitch:
		ce.Pitch = transformer.TransformValue(ce.Pitch)
	case TransformVelocity:
		ce.Velocity = transformer.TransformValue(ce.Velocity)
	}

	return ce
}

// Pause stream event
type Pause struct {
	basicEvent
}

// NewPause creates an initialized pause
func NewPause(duration float64) *Pause {
	return &Pause{
		basicEvent: basicEvent{
			duration:  duration,
			eventType: PauseEvent,
		},
	}
}

// String representation
func (e *Pause) String() string {
	return fmt.Sprintf("{duration: %v}", e.duration)
}

// ApplyTransform to pause
func (e *Pause) ApplyTransform(transformer Transformer) StreamEvent {
	ce := NewPause(e.duration)

	switch transformer.TransformType() {
	case TransformDuration:
		ce.duration = transformer.TransformValue(Value{ce.duration})[0]
	}

	return ce
}

// Streamer is an object that can produce a stream
type Streamer interface {
	Stream() *Stream
}

// Stream of events
type Stream struct {
	Events        []StreamEvent
	Duration      float64
	NumNoteEvents int
}

// NewStream creates an initialized stream
func NewStream() *Stream {
	return &Stream{
		Events: []StreamEvent{},
	}
}

// AddEvent to stream
func (s *Stream) AddEvent(event StreamEvent) {
	s.Events = append(s.Events, event)
	s.Duration += event.Duration()

	if event.Type() == NoteEvent {
		s.NumNoteEvents++
	}
}

// Append a stream and return a new stream
func (s *Stream) Append(sc *Stream) *Stream {
	copy := NewStream()

	for _, e := range s.Events {
		copy.AddEvent(e)
	}

	for _, e := range sc.Events {
		copy.AddEvent(e)
	}

	return copy
}

// ApplyTransform to stream
func (s *Stream) ApplyTransform(t Transformer) *Stream {
	return s.ApplyTransforms([]Transformer{t})
}

// ApplyTransforms to stream
func (s *Stream) ApplyTransforms(ts []Transformer) *Stream {
	cs := NewStream()

	for _, e := range s.Events {
		ce := e

		for _, t := range ts {
			ce = ce.ApplyTransform(t)
		}

		cs.AddEvent(ce)
	}

	return cs
}

// Sanitize the stream, combine consecutive pause events
func (s *Stream) Sanitize() {
	newEvents := []StreamEvent{}

	var previousPause *Pause

	for _, e := range s.Events {
		if currentPause, ok := e.(*Pause); ok {
			if previousPause != nil {
				previousPause.duration += currentPause.duration
			} else {
				previousPause = currentPause
			}
		} else if previousPause != nil {
			newEvents = append(newEvents, previousPause, e)
			previousPause = nil
		} else {
			newEvents = append(newEvents, e)
		}
	}

	if previousPause != nil {
		newEvents = append(newEvents, previousPause)
	}

	s.Events = newEvents
}

// TimedNote timed note representation to allow for easy
// conversion to midi notes. Also allows for easy combination of
// sections in one track
type TimedNote struct {
	Time      float64
	DeltaTime float64
	NoteOn    bool
	Pitch     float64
	Velocity  float64
	Channel   int
}

// TimedNotes convenience type
type TimedNotes []*TimedNote

// CalculateDeltaTimes for all timed notes
func (tns TimedNotes) CalculateDeltaTimes() {
	for index, rep := range tns {
		if index > 0 {
			rep.DeltaTime = rep.Time - tns[index-1].Time
		} else {
			rep.DeltaTime = 0
		}
	}
}

// Len for sorting
func (tns TimedNotes) Len() int {
	return len(tns)
}

// Swap for sorting
func (tns TimedNotes) Swap(i, j int) {
	tns[i], tns[j] = tns[j], tns[i]
}

// Less for sorting
func (tns TimedNotes) Less(i, j int) bool {
	return tns[i].Time < tns[j].Time
}

// String representation
func (tn *TimedNote) String() string {
	if tn.NoteOn {
		return fmt.Sprintf("{Note on - time: %v, pitch: %v, velocity: %v, channel: %v}", tn.Time, tn.Pitch, tn.Velocity, tn.Channel)
	}

	return fmt.Sprintf("{Note off - time: %v, pitch: %v, channel: %v}", tn.Time, tn.Pitch, tn.Channel)
}

// ToMidiTrack converts a sequence of timed notes to a midi track
func (tns TimedNotes) ToMidiTrack(ticksPerQuarterNote float64) *midi.Track {
	track := &midi.Track{
		Events: []midi.Event{},
	}

	var me *midi.ChannelEvent

	for _, note := range tns {
		me = &midi.ChannelEvent{}
		me.Channel = uint16(note.Channel)
		me.SetDeltaTime(uint32(note.DeltaTime * ticksPerQuarterNote))

		me.Value1 = uint16(note.Pitch)

		if note.NoteOn {
			me.SetEventType(midi.NoteOn)
			me.Value2 = uint16(note.Velocity)
		} else {
			me.SetEventType(midi.NoteOff)
			me.Value2 = 0
		}

		track.Events = append(track.Events, me)
	}

	endEvent := &midi.MetaEvent{
		MetaType: midi.EndOfTrack,
		Data:     []byte{},
	}

	track.Events = append(track.Events, endEvent)

	return track
}

// TimedNotes get timed note representation
func (s *Stream) TimedNotes(startTime float64) TimedNotes {
	time := startTime
	reps := TimedNotes{}

	for _, e := range s.Events {
		if e.Type() == NoteEvent {
			n := e.(*Note)
			// Add note on for all pitches
			for index, pitch := range n.Pitch {
				velocityIndex := int(math.Min(float64(index), float64(len(n.Velocity)-1)))
				rep := &TimedNote{
					Time:     time,
					NoteOn:   true,
					Pitch:    pitch,
					Velocity: n.Velocity[velocityIndex],
					Channel:  n.Channel,
				}

				reps = append(reps, rep)
			}

			time += e.Duration()

			// Add note off for all pitches
			for _, pitch := range n.Pitch {
				rep := &TimedNote{
					Time:    time,
					NoteOn:  false,
					Pitch:   pitch,
					Channel: n.Channel,
				}

				reps = append(reps, rep)
			}
		} else {
			time += e.Duration()
		}
	}

	return reps
}
