package export

import (
	"os"

	kallos "github.com/almerlucke/gokallos"
	midi "github.com/almerlucke/gomidi"
)

// StreamToMidiTrack converts a Kallos stream to a midi track
func StreamToMidiTrack(stream *kallos.Stream, ticksPerSecond float64) *midi.Track {
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
			channel := uint16(note.Channel)
			first := true

			for i, pitch := range note.Pitch {
				me = &midi.ChannelEvent{}

				var velocity float64

				if i >= len(note.Velocity) {
					velocity = note.Velocity[len(note.Velocity)-1]
				} else {
					velocity = note.Velocity[i]
				}

				if first {
					first = false
					me.SetDeltaTime(deltaTime)
				} else {
					me.SetDeltaTime(0)
				}

				me.SetEventType(midi.NoteOn)

				me.Channel = channel
				me.Value1 = uint16(pitch)
				me.Value2 = uint16(velocity)

				track.Events = append(track.Events, me)
			}

			deltaTime = uint32(note.Duration() * ticksPerSecond)
			first = true

			for _, pitch := range note.Pitch {
				me = &midi.ChannelEvent{}

				if first {
					first = false
					me.SetDeltaTime(deltaTime)
				} else {
					me.SetDeltaTime(0)
				}

				me.SetEventType(midi.NoteOff)
				me.Channel = channel
				me.Value1 = uint16(pitch)
				me.Value2 = uint16(0)

				track.Events = append(track.Events, me)
			}

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
		track := StreamToMidiTrack(s, float64(ticksPerBeat*2))
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
