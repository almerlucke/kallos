package generators

import (
	"testing"

	kallos "github.com/almerlucke/gokallos"
)

func chain1() *MarkovChain {
	state1 := NewMarkovState(NewSequence([]kallos.Value{60, 69, 64, 67, 65}, true), []float64{0.8, 0.1, 0.1}, nil)
	state2 := NewMarkovState(NewSequence([]kallos.Value{84, 93, 93, 86}, true), []float64{0.8, 0.1, 0.1}, nil)
	state3 := NewMarkovState(NewSequence([]kallos.Value{72, 65, 74, 83}, true), []float64{0.8, 0.1, 0.1}, nil)
	state4 := NewMarkovState(NewSequence([]kallos.Value{48, 60, 48, 52, 54}, true), []float64{0.8, 0.1, 0.1}, nil)

	state1.States = []*MarkovState{state2, state3, state4}
	state2.States = []*MarkovState{state3, state4, state1}
	state3.States = []*MarkovState{state4, state1, state2}
	state4.States = []*MarkovState{}

	return NewMarkovChain([]*MarkovState{state1, state2, state3, state4}, state1, false)
}

func chain2() *MarkovChain {
	state1 := NewMarkovState(NewSequence([]kallos.Value{62, 65, 64, 69, 69}, true), []float64{0.8, 0.1, 0.1}, nil)
	state2 := NewMarkovState(NewSequence([]kallos.Value{86, 89, 93, 86}, true), []float64{0.8, 0.1, 0.1}, nil)
	state3 := NewMarkovState(NewSequence([]kallos.Value{75, 67, 76, 83}, true), []float64{0.8, 0.1, 0.1}, nil)
	state4 := NewMarkovState(NewSequence([]kallos.Value{36, 60, 48, 36, 52}, true), []float64{0.8, 0.1, 0.1}, nil)

	state1.States = []*MarkovState{state2, state3, state4}
	state2.States = []*MarkovState{state3, state4, state1}
	state3.States = []*MarkovState{state4, state1, state2}
	state4.States = []*MarkovState{}

	return NewMarkovChain([]*MarkovState{state1, state2, state3, state4}, state1, false)
}

func chain3() *MarkovChain {
	state1 := NewMarkovState(NewSequence([]kallos.Value{60, 65, 60, 69, 69}, true), []float64{0.8, 0.1, 0.1}, nil)
	state2 := NewMarkovState(NewSequence([]kallos.Value{84, 89, 84, 86}, true), []float64{0.8, 0.1, 0.1}, nil)
	state3 := NewMarkovState(NewSequence([]kallos.Value{72, 67, 72, 83}, true), []float64{0.8, 0.1, 0.1}, nil)
	state4 := NewMarkovState(NewSequence([]kallos.Value{36, 60, 48, 36, 52}, true), []float64{0.8, 0.1, 0.1}, nil)

	state1.States = []*MarkovState{state2, state3, state4}
	state2.States = []*MarkovState{state3, state4, state1}
	state3.States = []*MarkovState{state4, state1, state2}
	state4.States = []*MarkovState{}

	return NewMarkovChain([]*MarkovState{state1, state2, state3, state4}, state1, false)
}

func TestMarkov(t *testing.T) {
	state1 := NewMarkovState(chain1(), []float64{0.7, 0.3}, nil)
	state2 := NewMarkovState(chain2(), []float64{0.7, 0.3}, nil)
	state3 := NewMarkovState(chain3(), []float64{0.7, 0.3}, nil)

	state1.States = []*MarkovState{state2, state3}
	state2.States = []*MarkovState{state3, state1}
	state3.States = []*MarkovState{state1, state2}

	chain := NewMarkovChain([]*MarkovState{state1, state2, state3}, state1, true)

	i := 0
	for {
		if i > 40 {
			break
		}

		t.Logf("markov value: %v", chain.GenerateValue())
		i++
	}
}
