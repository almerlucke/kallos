package tests

import (
	"testing"

	"github.com/almerlucke/kallos"
	gens "github.com/almerlucke/kallos/generators"
)

func chain1() *gens.MarkovChain {
	state1 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(60, 69, 64, 67, 65), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)
	state2 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(84, 93, 93, 86), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)
	state3 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(72, 65, 74, 83), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)
	state4 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(48, 60, 48, 52, 54), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)

	state1.States = []*gens.MarkovState{state2, state3, state4}
	state2.States = []*gens.MarkovState{state3, state4, state1}
	state3.States = []*gens.MarkovState{state4, state1, state2}
	state4.States = []*gens.MarkovState{}

	return gens.NewMarkovChain([]*gens.MarkovState{state1, state2, state3, state4}, state1, false)
}

func chain2() *gens.MarkovChain {
	state1 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(62, 65, 64, 69, 69), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)
	state2 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(86, 89, 93, 86), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)
	state3 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(75, 67, 76, 83), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)
	state4 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(36, 60, 48, 36, 52), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)

	state1.States = []*gens.MarkovState{state2, state3, state4}
	state2.States = []*gens.MarkovState{state3, state4, state1}
	state3.States = []*gens.MarkovState{state4, state1, state2}
	state4.States = []*gens.MarkovState{}

	return gens.NewMarkovChain([]*gens.MarkovState{state1, state2, state3, state4}, state1, false)
}

func chain3() *gens.MarkovChain {
	state1 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(60, 65, 60, 69, 69), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)
	state2 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(84, 89, 84, 86), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)
	state3 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(72, 67, 72, 83), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)
	state4 := gens.NewMarkovState(gens.NewSequence(kallos.ToValues(36, 60, 48, 36, 52), true), gens.Probabilities{0.8, 0.1, 0.1}, nil)

	state1.States = []*gens.MarkovState{state2, state3, state4}
	state2.States = []*gens.MarkovState{state3, state4, state1}
	state3.States = []*gens.MarkovState{state4, state1, state2}
	state4.States = []*gens.MarkovState{}

	return gens.NewMarkovChain([]*gens.MarkovState{state1, state2, state3, state4}, state1, false)
}

func TestMarkov(t *testing.T) {
	state1 := gens.NewMarkovState(chain1(), gens.Probabilities{0.7, 0.3}, nil)
	state2 := gens.NewMarkovState(chain2(), gens.Probabilities{0.7, 0.3}, nil)
	state3 := gens.NewMarkovState(chain3(), gens.Probabilities{0.7, 0.3}, nil)

	state1.States = []*gens.MarkovState{state2, state3}
	state2.States = []*gens.MarkovState{state3, state1}
	state3.States = []*gens.MarkovState{state1, state2}

	chain := gens.NewMarkovChain([]*gens.MarkovState{state1, state2, state3}, state1, true)

	i := 0
	for {
		if i > 40 {
			break
		}

		t.Logf("markov value: %v", chain.GenerateValue())
		i++
	}
}
