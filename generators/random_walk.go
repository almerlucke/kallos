package generators

import (
	"math/rand"

	"github.com/almerlucke/kallos"
)

type RandomWalkMatrix interface {
	ValueForIndex(index []int) kallos.Value
}

type RandomWalk1DMatrix struct {
	Values []kallos.Value
}

func (r *RandomWalk1DMatrix) ValueForIndex(index []int) kallos.Value {
	return r.Values[index[0]]
}

type RandomWalk2DMatrix struct {
	Values [][]kallos.Value
}

func (r *RandomWalk2DMatrix) ValueForIndex(index []int) kallos.Value {
	return r.Values[index[0]][index[1]]
}

type RandomWalk3DMatrix struct {
	Values [][][]kallos.Value
}

func (r *RandomWalk3DMatrix) ValueForIndex(index []int) kallos.Value {
	return r.Values[index[0]][index[1]][index[2]]
}

type RandomWalk struct {
	Dimensions []int
	Index      []int
	Matrix     RandomWalkMatrix
}

func NewRandomWalk(dimensions []int, matrix RandomWalkMatrix) *RandomWalk {
	rw := &RandomWalk{
		Dimensions: dimensions,
		Matrix:     matrix,
	}

	index := make([]int, len(dimensions))

	for i, d := range dimensions {
		index[i] = rand.Intn(d)
	}

	rw.Index = index

	return rw
}

func (r *RandomWalk) GenerateValue() kallos.Value {
	selectedDimension := rand.Intn(len(r.Dimensions))
	selectedDimensionSize := r.Dimensions[selectedDimension]
	selectedIndex := r.Index[selectedDimension]

	if selectedIndex < 1 {
		selectedIndex++
	} else if selectedIndex > (selectedDimensionSize - 2) {
		selectedIndex--
	} else {
		if rand.Intn(2) == 1 {
			selectedIndex++
		} else {
			selectedIndex--
		}
	}

	r.Index[selectedDimension] = selectedIndex

	return r.Matrix.ValueForIndex(r.Index)
}

func (r *RandomWalk) IsContinuous() bool {
	return true
}

func (r *RandomWalk) Done() bool {
	return false
}

func (r *RandomWalk) Reset() {
}
