package gokallos

import (
	"fmt"
	midi "github.com/almerlucke/gomidi"
)

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
func (tns TimedNotes) ToMidiTrack(ticksPerSecond float64) *midi.Track {
	track := &midi.Track{
		Events: []midi.Event{},
	}

	var me *midi.ChannelEvent

	for _, note := range tns {
		me = &midi.ChannelEvent{}
		me.Channel = uint16(note.Channel)
		me.SetDeltaTime(uint32(note.DeltaTime * ticksPerSecond))

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
