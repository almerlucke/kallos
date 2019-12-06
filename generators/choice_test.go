package generators

import (
	"math/rand"
	"testing"

	kallos "github.com/almerlucke/gokallos"
)

func TestCore(t *testing.T) {

	// seed := time.Now().UTC().UnixNano()

	rand.Seed(12232)

	s := &kallos.Section{}
	s.Clock = 0.01
	s.Until = &kallos.LengthStopCondition{
		Length: 10,
	}
	s.Rhythm = &RandomChoice{
		Elements: []kallos.Value{-0.25, 0.25, 1.0, -1.0, 2.0, 1.5},
	}
	s.Pitch = &RandomChoice{
		Elements: []kallos.Value{60.0, 63.0, 65.0, 66.0},
	}
	s.Velocity = &RandomChoice{
		Elements: []kallos.Value{100.0, 120.0, 60.0, 30.0},
	}
	s.Channel = &RandomChoice{
		Elements: []kallos.Value{1},
	}

	stream := s.Stream()

	for _, e := range stream.Events {
		t.Logf("event %v", e)
	}
}
