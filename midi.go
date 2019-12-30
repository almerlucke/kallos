package kallos

import (
	"os"

	midi "github.com/almerlucke/gomidi"
)

// MidiTracker type that can generate a midi track
type MidiTracker interface {
	ToMidiTrack() *midi.Track
}

// ToMidiFile create file from midi trackers
func ToMidiFile(filePath string, trackers []MidiTracker) error {
	header := midi.FileHeader{}
	header.Format = midi.Format1
	header.NumTracks = uint16(len(trackers))
	header.Division = uint16(TicksPerQuarterNote)

	midiFile := midi.NewFile()
	midiFile.Chunks = append(midiFile.Chunks, header.Chunk())

	for _, tracker := range trackers {
		track := tracker.ToMidiTrack()
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
