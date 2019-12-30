package gokallos

import (
	"os"

	midi "github.com/almerlucke/gomidi"
)

// MidiTracker type that can generate a midi track
type MidiTracker interface {
	ToMidiTrack(ticksPerQuarterNote float64) *midi.Track
}

// MidiTrackersToFile create file from midi track generators
func MidiTrackersToFile(trackers []MidiTracker, filePath string) error {
	ticksPerQuarterNote := 192.0

	header := midi.FileHeader{}
	header.Format = midi.Format1
	header.NumTracks = uint16(len(trackers))
	header.Division = uint16(ticksPerQuarterNote)

	midiFile := midi.NewFile()
	midiFile.Chunks = append(midiFile.Chunks, header.Chunk())

	for _, tracker := range trackers {
		track := tracker.ToMidiTrack(ticksPerQuarterNote)
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
