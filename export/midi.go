package export

import (
	"os"

	kallos "github.com/almerlucke/gokallos"
	midi "github.com/almerlucke/gomidi"
)

// StreamToMidiTrack converts a Kallos stream to a midi track
func StreamToMidiTrack(stream *kallos.Stream, ticksPerSecond kallos.Value) *midi.Track {
	track := &midi.Track{
		Events: []midi.Event{},
	}

	var me *midi.ChannelEvent

	deltaTime := uint32(0)
	for _, se := range stream.Events {
		if pause, ok := se.(*kallos.Pause); ok {
			deltaTime = uint32(pause.Duration() * ticksPerSecond)
		} else {
			note := se.(*kallos.Note)

			me = &midi.ChannelEvent{}
			me.SetDeltaTime(deltaTime)
			me.SetEventType(midi.NoteOn)

			channel := uint16(kallos.Clip(note.Channel, kallos.Value(0.0), kallos.Value(15.0)))
			pitch := uint16(kallos.Clip(note.Pitch, kallos.Value(0.0), kallos.Value(127.0)))
			me.Channel = channel
			me.Value1 = pitch
			me.Value2 = uint16(kallos.Clip(note.Velocity, kallos.Value(0.0), kallos.Value(127.0)))

			track.Events = append(track.Events, me)

			deltaTime = uint32(note.Duration() * ticksPerSecond)
			me = &midi.ChannelEvent{}
			me.SetDeltaTime(deltaTime)
			me.SetEventType(midi.NoteOff)
			me.Channel = channel
			me.Value1 = pitch
			me.Value2 = uint16(0)

			track.Events = append(track.Events, me)

			deltaTime = 0
		}
	}

	endEvent := &midi.MetaEvent{
		MetaType: midi.EndOfTrack,
		Data:     []byte{},
	}

	track.Events = append(track.Events, endEvent)

	return track
}

// StreamsToMidiFile writes a sequence of streams as separate tracks to a midi file
func StreamsToMidiFile(streams []*kallos.Stream, ticksPerBeat uint16, filePath string) error {

	header := midi.FileHeader{}
	header.Format = midi.Format1
	header.NumTracks = uint16(len(streams))
	header.Division = ticksPerBeat

	midiFile := midi.NewFile()
	midiFile.Chunks = append(midiFile.Chunks, header.Chunk())

	for _, s := range streams {
		track := StreamToMidiTrack(s, kallos.Value(ticksPerBeat*2))
		midiFile.Chunks = append(midiFile.Chunks, track.Chunk())
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	midiFile.WriteTo(f)

	return nil
}
