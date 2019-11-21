package export

import (
	"math/rand"
	"testing"
	"time"

	kallos "github.com/almerlucke/gokallos"
	"github.com/almerlucke/gokallos/generators"
)

func TestExportMidi(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	rand.Seed(seed)

	s := &kallos.Section{}
	s.Clock = 1.0
	s.Until = &kallos.LengthStopCondition{
		Length: 100,
	}
	s.Rhythm = &generators.RandomChoice{
		Elements: []kallos.Value{-0.125, 0.125, 0.25, 0.25, 0.5, 0.25, 0.25, 0.125},
	}
	s.Pitch = &generators.RandomChoice{
		Elements: []kallos.Value{60.0, 64.0, 65.0, 69.0, 72.0, 76.0, 77.0, 81.0},
	}
	s.Velocity = &generators.RandomChoice{
		Elements: []kallos.Value{100.0, 120.0, 80.0, 90.0},
	}
	s.Channel = &generators.RandomChoice{
		Elements: []kallos.Value{1},
	}

	stream := s.Stream()
	stream.Sanitize()

	streams := []*kallos.Stream{stream}

	s.Until = &kallos.LengthStopCondition{
		Length: 70,
	}
	s.Rhythm = &generators.RandomChoice{
		Elements: []kallos.Value{0.25, 0.25, 0.5, 0.5, 1.0, 0.25, 0.125, -0.25},
	}
	s.Pitch = &generators.RandomChoice{
		Elements: []kallos.Value{48.0, 52.0, 53.0, 57.0, 60.0, 64.0, 65.0, 69.0},
	}
	s.Velocity = &generators.RandomChoice{
		Elements: []kallos.Value{100.0, 120.0, 80.0, 90.0},
	}
	s.Channel = &generators.RandomChoice{
		Elements: []kallos.Value{1},
	}

	stream = s.Stream()
	stream.Sanitize()

	streams = append(streams, stream)

	StreamsToMidiFile(streams, uint16(96), "test_output.mid")
}
