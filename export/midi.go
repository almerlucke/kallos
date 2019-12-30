package export

import (
	"os"

	kallos "github.com/almerlucke/gokallos"
	midi "github.com/almerlucke/gomidi"
)

// StreamToMidiTrack converts a Kallos stream to a midi track
func StreamToMidiTrack(stream *kallos.Stream, ticksPerQuarterNote float64) *midi.Track {
	track := &midi.Track{
		Events: []midi.Event{},
	}

	var me *midi.ChannelEvent

	notes := stream.TimedNotes(0)

	notes.CalculateDeltaTimes()

	for _, note := range notes {
		me = &midi.ChannelEvent{}
		me.Channel = uint16(note.Channel)
		me.SetDeltaTime(uint32(note.DeltaTime * ticksPerQuarterNote * 4))

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

// TracksToMidiFile tracks to midi file
func TracksToMidiFile(tracks []*midi.Track, ticksPerQuarterNote uint16, filePath string) error {
	header := midi.FileHeader{}
	header.Format = midi.Format1
	header.NumTracks = uint16(len(tracks))
	header.Division = ticksPerQuarterNote

	midiFile := midi.NewFile()
	midiFile.Chunks = append(midiFile.Chunks, header.Chunk())

	for _, track := range tracks {
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

// StreamsToMidiFile writes a sequence of streams as separate tracks to a midi file
func StreamsToMidiFile(streams []*kallos.Stream, ticksPerQuarterNote uint16, filePath string) error {

	header := midi.FileHeader{}
	header.Format = midi.Format1
	header.NumTracks = uint16(len(streams))
	header.Division = ticksPerQuarterNote

	midiFile := midi.NewFile()
	midiFile.Chunks = append(midiFile.Chunks, header.Chunk())

	for _, s := range streams {
		track := StreamToMidiTrack(s, float64(ticksPerQuarterNote))
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
