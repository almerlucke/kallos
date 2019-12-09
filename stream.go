package gokallos

import (
	"fmt"
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
	Events   []StreamEvent
	Duration float64
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
