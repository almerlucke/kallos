package tests

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/almerlucke/kallos"

	"github.com/almerlucke/kallos/generators"
	"github.com/almerlucke/kallos/generators/rhythm"
	"github.com/almerlucke/kallos/generators/tools"
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

	// streams := []*Stream{}

	// matrix := &generators.RandomWalk2DMatrix{
	// 	Values: [][]Value{
	// 		kallos.ToValues(36, 38, 39, 42, 43, 44),
	// 		kallos.ToValues(47, 48, 50, 51, 52, 54),
	// 		kallos.ToValues(55, 58, 59, 60, 63, 64),
	// 		kallos.ToValues(66, 65, 67, 69, 70, 73),
	// 		kallos.ToValues(74, 75, 78, 79, 81, 82),
	// 		kallos.ToValues(84, 86, 87, 90, 91, 92),
	// 	},
	// }

	//walker := generators.NewRandomWalk([]int{6, 6}, matrix)

	shape := kallos.CreateShape(func(i int, n int) float64 {
		return math.Sin(math.Pi*2*float64(i)/float64(n))*0.5 + 0.5
	}, 128)

	velocities := kallos.ToValues(30, 40, 50, 60, 70, 80, 90, 100, 110, 120)

	note, _ := kallos.NoteNumberFromNoteName("c4")

	arpeggio1 := generators.NewArpeggio(note, []int{0, 4, 5, 9, -2, -3, -5, -7}, []int{3, 2, 1, 0, 2, 1, 0}, true)
	// arpeggio2 := generators.NewArpeggio(note, []int{0, -2, -3, -5, -7, 0, 2, 5, 7}, []int{2, 1, 0, 3}, true)

	// pitchChain := pitchCombinedChain()

	rhythm := kallos.NewRhythm(
		0.5,
		kallos.NewDurationStopCondition(30.0),
		rhythm.NewBouncer(
			tools.NewRamp(10, 0.125, 0.5, 0.6),
			tools.NewRamp(10, 0.125, 0.5, 0.6),
			tools.NewRamp(6, 0.25, 1.0, 0.8),
		),
	)

	s1 := &kallos.Section{}
	s1.Clock = 0.5
	s1.Until = kallos.NewLengthStopCondition(uint32(rhythm.NumNoteEvents()))
	s1.Rhythm = generators.NewSequence(rhythm.Values().Apply2(kallos.Round, 0.125), true)
	s1.Pitch = arpeggio1
	s1.Velocity = generators.NewSequence(shape.Convert(velocities, rhythm.NumNoteEvents()).Apply(math.Round), true)
	s1.Channel = generators.NewStaticValue(kallos.Value{1})

	// s2 := &kallos.Section{}
	// s2.Clock = 0.5
	// s2.Until = &kallos.LengthStopCondition{
	// 	Length: 30,
	// }
	// s2.Rhythm = rhythm.NewBouncer(
	// 	tools.NewRamp(10, 0.125, 0.5, 0.6),
	// 	tools.NewRamp(10, 0.125, 0.5, 0.6),
	// 	tools.NewRamp(6, 0.25, 1.0, 0.8),
	// )
	// s2.Pitch = arpeggio2
	// s2.Velocity = generators.NewRamp(tools.NewRamp(10, 110, 20, 0.6), true)
	// s2.Channel = generators.NewStaticValue(kallos.Value{1})

	// ts := kallos.SequentialSection{
	// 	s1,
	// 	s2,
	// }

	kallos.ToMidiFile("test_output.mid", []kallos.MidiTracker{
		s1,
	})
}
