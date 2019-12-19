package export

import (
	"math/rand"
	"testing"
	"time"

	kallos "github.com/almerlucke/gokallos"
	"github.com/almerlucke/gokallos/generators"
)

func pitchChain1() *generators.MarkovChain {
	state1 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(60, 62, 63, 65, 67), false), []float64{0.8, 0.1, 0.1}, nil)
	state2 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(72, 74, 75, 74), false), []float64{0.8, 0.1, 0.1}, nil)
	state3 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(72, 75, 77, 75), false), []float64{0.8, 0.1, 0.1}, nil)
	state4 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(77, 78, 80, 82, 83), false), []float64{0.8, 0.1, 0.1}, nil)

	state1.States = []*generators.MarkovState{state2, state3, state4}
	state2.States = []*generators.MarkovState{state3, state4, state1}
	state3.States = []*generators.MarkovState{state4, state1, state2}
	state4.States = []*generators.MarkovState{}

	return generators.NewMarkovChain([]*generators.MarkovState{state1, state2, state3, state4}, state1, false)
}

func pitchChain2() *generators.MarkovChain {
	state1 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(60, 63, 62, 67, 65), false), []float64{0.8, 0.1, 0.1}, nil)
	state2 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(74, 72, 74, 75), false), []float64{0.8, 0.1, 0.1}, nil)
	state3 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(75, 72, 75, 77), false), []float64{0.8, 0.1, 0.1}, nil)
	state4 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(83, 82, 80, 78, 77), false), []float64{0.8, 0.1, 0.1}, nil)

	state1.States = []*generators.MarkovState{state2, state3, state4}
	state2.States = []*generators.MarkovState{state3, state4, state1}
	state3.States = []*generators.MarkovState{state4, state1, state2}
	state4.States = []*generators.MarkovState{}

	return generators.NewMarkovChain([]*generators.MarkovState{state1, state2, state3, state4}, state1, false)
}

func pitchChain3() *generators.MarkovChain {
	state1 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(67, 67, 67, 67, 65), false), []float64{0.8, 0.1, 0.1}, nil)
	state2 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(72, 72, 72, 75), false), []float64{0.8, 0.1, 0.1}, nil)
	state3 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(77, 75, 72, 75), false), []float64{0.8, 0.1, 0.1}, nil)
	state4 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(82, 83, 78, 80, 77), false), []float64{0.8, 0.1, 0.1}, nil)

	state1.States = []*generators.MarkovState{state2, state3, state4}
	state2.States = []*generators.MarkovState{state3, state4, state1}
	state3.States = []*generators.MarkovState{state4, state1, state2}
	state4.States = []*generators.MarkovState{}

	return generators.NewMarkovChain([]*generators.MarkovState{state1, state2, state3, state4}, state1, false)
}

func pitchCombinedChain() *generators.MarkovChain {
	state1 := generators.NewMarkovState(pitchChain1(), []float64{0.7, 0.3}, nil)
	state2 := generators.NewMarkovState(pitchChain2(), []float64{0.7, 0.3}, nil)
	state3 := generators.NewMarkovState(pitchChain3(), []float64{0.7, 0.3}, nil)

	state1.States = []*generators.MarkovState{state2, state3}
	state2.States = []*generators.MarkovState{state3, state1}
	state3.States = []*generators.MarkovState{state1, state2}

	return generators.NewMarkovChain([]*generators.MarkovState{state1, state2, state3}, state1, true)
}

func rhythmChain1() *generators.MarkovChain {
	state1 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.125, 0.125, 0.125, 0.25, 0.375), false), []float64{0.8, 0.1, 0.1}, nil)
	state2 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.25, 0.25, 0.25, 0.75), false), []float64{0.8, 0.1, 0.1}, nil)
	state3 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.125, 0.25, 0.125, 0.25), false), []float64{0.8, 0.1, 0.1}, nil)
	state4 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.375, 0.375, 0.25, 0.125, 0.125), false), []float64{0.8, 0.1, 0.1}, nil)

	state1.States = []*generators.MarkovState{state2, state3, state4}
	state2.States = []*generators.MarkovState{state3, state4, state1}
	state3.States = []*generators.MarkovState{state4, state1, state2}
	state4.States = []*generators.MarkovState{}

	return generators.NewMarkovChain([]*generators.MarkovState{state1, state2, state3, state4}, state1, false)
}

func rhythmChain2() *generators.MarkovChain {
	state1 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.125, 0.125, 0.125, 0.25, 0.375), false), []float64{0.8, 0.1, 0.1}, nil)
	state2 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.75, 0.25, 0.25, 0.25), false), []float64{0.8, 0.1, 0.1}, nil)
	state3 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.25, 0.125, 0.25, 0.125), false), []float64{0.8, 0.1, 0.1}, nil)
	state4 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.125, 0.125, 0.25, 0.375, 0.375), false), []float64{0.8, 0.1, 0.1}, nil)

	state1.States = []*generators.MarkovState{state2, state3, state4}
	state2.States = []*generators.MarkovState{state3, state4, state1}
	state3.States = []*generators.MarkovState{state4, state1, state2}
	state4.States = []*generators.MarkovState{}

	return generators.NewMarkovChain([]*generators.MarkovState{state1, state2, state3, state4}, state1, false)
}

func rhythmChain3() *generators.MarkovChain {
	state1 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.125, 0.25, 0.375, 0.125, 0.125), false), []float64{0.8, 0.1, 0.1}, nil)
	state2 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.75, 0.25, 0.25, 0.25), false), []float64{0.8, 0.1, 0.1}, nil)
	state3 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.25, 0.125, 0.25, 0.125), false), []float64{0.8, 0.1, 0.1}, nil)
	state4 := generators.NewMarkovState(generators.NewSequence(kallos.ToValues(0.125, 0.125, 0.25, 0.375, 0.375), false), []float64{0.8, 0.1, 0.1}, nil)

	state1.States = []*generators.MarkovState{state2, state3, state4}
	state2.States = []*generators.MarkovState{state3, state4, state1}
	state3.States = []*generators.MarkovState{state4, state1, state2}
	state4.States = []*generators.MarkovState{}

	return generators.NewMarkovChain([]*generators.MarkovState{state1, state2, state3, state4}, state1, false)
}

func rhythmCombinedChain() *generators.MarkovChain {
	state1 := generators.NewMarkovState(rhythmChain1(), []float64{0.7, 0.3}, nil)
	state2 := generators.NewMarkovState(rhythmChain2(), []float64{0.7, 0.3}, nil)
	state3 := generators.NewMarkovState(rhythmChain3(), []float64{0.7, 0.3}, nil)

	state1.States = []*generators.MarkovState{state2, state3}
	state2.States = []*generators.MarkovState{state3, state1}
	state3.States = []*generators.MarkovState{state1, state2}

	return generators.NewMarkovChain([]*generators.MarkovState{state1, state2, state3}, state1, true)
}

func TestExportMidi(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	rand.Seed(seed)

	streams := []*kallos.Stream{}

	matrix := &generators.RandomWalk2DMatrix{
		Values: [][]kallos.Value{
			kallos.ToValues(36, 38, 39, 42, 43, 44),
			kallos.ToValues(47, 48, 50, 51, 52, 54),
			kallos.ToValues(55, 58, 59, 60, 63, 64),
			kallos.ToValues(66, 65, 67, 69, 70, 73),
			kallos.ToValues(74, 75, 78, 79, 81, 82),
			kallos.ToValues(84, 86, 87, 90, 91, 92),
		},
	}

	walker := generators.NewRandomWalk([]int{6, 6}, matrix)

	// note, _ := kallos.NoteNumberFromNoteName("c4")

	// arpeggio1 := generators.NewArpeggio(note, []int{0, 4, 5, 9, -2, -3, -5, -7}, []int{0, 1, 2, 1}, true)
	// arpeggio2 := generators.NewArpeggio(note, []int{0, -2, -3, -5, -7, -8, 0, 2, 5, 7}, []int{2, 1, 0, 1}, true)
	// sequence := generators.NewSequence(kallos.ToValues(36, 36, 38, 38, 40, 40, 41, 41, 40, 40), true)
	// combinator := generators.NewCombinator(arpeggio1, arpeggio2, sequence)

	// pitchChain := pitchCombinedChain()

	s := &kallos.Section{}
	s.Clock = 1.0
	s.Until = &kallos.LengthStopCondition{
		Length: 100,
	}
	s.Rhythm = rhythmCombinedChain()
	s.Pitch = walker
	s.Velocity = generators.NewRandomChoice(kallos.ToValues(100.0, 120.0, 80.0, 90.0), true, true)
	s.Channel = generators.NewStaticValue(kallos.Value{1})

	// stream := s.Stream()
	// transformer := transformers.NewMultiplier(2, kallos.TransformDuration)
	// stream2 := stream.ApplyTransform(transformer)

	streams = append(streams, s.Stream())

	StreamsToMidiFile(streams, uint16(96), "test_output.mid")
}
