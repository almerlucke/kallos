package exporters

import (
	"math"
	"math/rand"
	"os"
	"testing"
	"time"

	kallos "github.com/almerlucke/gokallos"
	"github.com/almerlucke/gokallos/generators"
	midi "github.com/almerlucke/gomidi"
)

func clip(v kallos.Value, low kallos.Value, high kallos.Value) kallos.Value {
	return kallos.Value(math.Min(math.Max(float64(v), float64(low)), float64(high)))
}

func streamToTrack(stream *kallos.Stream, ticksPerSecond kallos.Value) *midi.Track {
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

			channel := uint16(clip(note.Channel, kallos.Value(0.0), kallos.Value(15.0)))
			pitch := uint16(clip(note.Pitch, kallos.Value(0.0), kallos.Value(127.0)))

			me.Channel = channel
			me.Value1 = pitch
			me.Value2 = uint16(clip(note.Velocity, kallos.Value(0.0), kallos.Value(127.0)))
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

func TestExportMidi(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	rand.Seed(seed)

	s := &kallos.Section{}
	s.Clock = 1.0
	s.Until = &kallos.LengthStopCondition{
		Length: 100,
	}
	s.Rhythm = &generators.RandomChoice{
		Elements: []kallos.Value{-0.25, 0.25, 0.5, 0.5, 1.0, -1.0},
	}
	s.Pitch = &generators.RandomChoice{
		Elements: []kallos.Value{60.0, 64.0, 65.0, 69.0, 72.0, 76.0, 77.0, 81.0},
	}
	s.Velocity = &generators.RandomChoice{
		Elements: []kallos.Value{100.0, 120.0, 60.0, 30.0},
	}
	s.Channel = &generators.RandomChoice{
		Elements: []kallos.Value{1},
	}

	header := midi.FileHeader{}
	header.Format = midi.Format1
	header.NumTracks = 2
	header.Division = 192

	midiFile := midi.NewFile()
	midiFile.Chunks = append(midiFile.Chunks, header.Chunk())

	stream := s.Stream()
	stream.Sanitize()
	track := streamToTrack(stream, kallos.Value(192))
	midiFile.Chunks = append(midiFile.Chunks, track.Chunk())

	stream = s.Stream()
	stream.Sanitize()
	track = streamToTrack(stream, kallos.Value(192))
	midiFile.Chunks = append(midiFile.Chunks, track.Chunk())

	f, err := os.Create("test_output.mid")
	if err != nil {
		t.Logf("err %v", err)
		return
	}

	defer f.Close()

	midiFile.WriteTo(f)
}
